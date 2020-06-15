package network

type Listener interface {
	Handle(handler ConnHandler)
	BroadcastSend(path string, event Event, excludeConn Conn)
}

type ConnHandler interface {
	ServeConn(conn Conn)
}
