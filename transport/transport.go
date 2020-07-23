package transport

type Listener interface {
	Listen()
	OnConn(handler ConnHandler)
}

type ConnHandler interface {
	ServeConn(conn Conn)
}

type Conn interface {
	SendMessage(path string, message interface{})
	Close()
	OnClose(handler CloseHandler)
	OnMessage(path string, handler MessageHandler)
}

type MessageHandler interface {
	NewMessage() interface{}
	ServeMessage(message interface{})
}

type CloseHandler interface {
	ServeClose()
}
