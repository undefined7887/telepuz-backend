package network

type Listener interface {
	Handle(handler ConnHandler)
}

type ConnHandler interface {
	ServeConn(conn Conn)
}
