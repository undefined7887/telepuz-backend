package network

type Listener interface {
	Handle(handler ConnHandler)
	BroadcastSend(path string, event Event)
}

type ConnHandler func(conn Conn)
