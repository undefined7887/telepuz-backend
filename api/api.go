package api

import "github.com/undefined7887/telepuz-backend/base/network"

func NewApi(listener network.Listener) {
	listener.Handle(handleConn)
}

func handleConn(conn network.Conn) {

}
