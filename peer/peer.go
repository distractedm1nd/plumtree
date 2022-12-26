package peer

import (
	"fmt"
	logger "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"sync"
)

var log = logger.Logger("tracker")

const (
	DefaultFanout = 10
)

type Sampler interface {
	FindPeers(count int) []peer.ID
	NeighborDown(peer.ID)
	NeighborUp(peer.ID)
}

type Tracker struct {
	host    host.Host
	sampler Sampler

	eagerMu sync.RWMutex
	Eager   map[peer.ID]struct{}
	lazyMu  sync.RWMutex
	Lazy    map[peer.ID]struct{}
}

func NewTracker(host host.Host, sampler Sampler) *Tracker {
	return &Tracker{
		host:    host,
		sampler: sampler,
		Eager:   make(map[peer.ID]struct{}),
		Lazy:    make(map[peer.ID]struct{}),
	}
}

func (t *Tracker) Start() {
	log.Debug("starting node", "peerID", t.host.ID())
	eagerPeers := t.sampler.FindPeers(DefaultFanout)
	t.eagerMu.Lock()
	defer t.eagerMu.Unlock()
	for _, p := range eagerPeers {
		log.Debug("initializing with eager peer", "peer", p.String())
		t.Eager[p] = struct{}{}
	}
}

// OnGraft attempts to graft the given peer to the eager list
// for eager pushing, and removes the given peer from the lazy list
// if it exists.
func (t *Tracker) OnGraft(p peer.ID) error {
	return t.moveToEager(p)
}

// OnPrune attempts to prune the given peer from the eager list
// and adds the given peer to the lazy list if it doesn't
// already exist.
func (t *Tracker) OnPrune(p peer.ID) error {
	return t.moveToLazy(p)
}

// moveToEager moves the given peer to the eager peerset
func (t *Tracker) moveToEager(p peer.ID) error {
	log.Debug("moving peer to eager set", "peer", p.String())
	t.lazyMu.Lock()
	if _, ok := t.Lazy[p]; ok {
		delete(t.Lazy, p)
	}
	t.lazyMu.Unlock()

	t.eagerMu.Lock()
	defer t.eagerMu.Unlock()
	_, ok := t.Eager[p]
	if ok {
		return fmt.Errorf("peer %s already in eager peerset", p.String())
	}
	t.Eager[p] = struct{}{}
	return nil
}

// moveToLazy moves the given peer to the lazy peerset
func (t *Tracker) moveToLazy(p peer.ID) error {
	log.Debug("moving peer to lazy set", "peer", p.String())
	t.eagerMu.Lock()
	if _, ok := t.Eager[p]; ok {
		delete(t.Eager, p)
	}
	t.eagerMu.Unlock()

	t.lazyMu.Lock()
	defer t.lazyMu.Unlock()
	_, ok := t.Lazy[p]
	if ok {
		return fmt.Errorf("peer %s already in lazy peerset", p.String())
	}
	t.Lazy[p] = struct{}{}
	return nil
}


