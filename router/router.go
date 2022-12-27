package router

import (
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
)

var _ pubsub.PubSubRouter = (*Router)(nil)

type Router struct {
}

func (r *Router) Protocols() []protocol.ID {
	//TODO implement me
	panic("implement me")
}

func (r *Router) Attach(sub *pubsub.PubSub) {
	//TODO implement me
	panic("implement me")
}

func (r *Router) AddPeer(id peer.ID, id2 protocol.ID) {
	//TODO implement me
	panic("implement me")
}

func (r *Router) RemovePeer(id peer.ID) {
	//TODO implement me
	panic("implement me")
}

func (r *Router) EnoughPeers(topic string, suggested int) bool {
	//TODO implement me
	panic("implement me")
}

func (r *Router) AcceptFrom(id peer.ID) pubsub.AcceptStatus {
	//TODO implement me
	panic("implement me")
}

func (r *Router) HandleRPC(rpc *pubsub.RPC) {
	//TODO implement me
	panic("implement me")
}

func (r *Router) Publish(message *pubsub.Message) {
	//TODO implement me
	panic("implement me")
}

func (r *Router) Join(topic string) {
	//TODO implement me
	panic("implement me")
}

func (r *Router) Leave(topic string) {
	//TODO implement me
	panic("implement me")
}
