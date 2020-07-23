package main

import (
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/services/common"
	"github.com/undefined7887/telepuz-backend/services/messages"
	"github.com/undefined7887/telepuz-backend/services/users"
	"github.com/undefined7887/telepuz-backend/transport"
	"github.com/undefined7887/telepuz-backend/utils"
	"os"
	"regexp"
)

var addrRegexp = regexp.MustCompile("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]).){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]):[0-9]+$")

func main() {
	logger := log.NewDefaultLogger("")

	if len(os.Args) < 2 {
		logger.Fatal("You should pass listening address as parameter")
	}

	addr := os.Args[1]
	if !addrRegexp.MatchString(addr) {
		logger.Fatal("Wrong address format: address should be like ip:port")
	}

	clientPool := utils.NewPool()
	userPool := utils.NewPool()

	listener := transport.NewWebsocketListener(logger, addr)
	listener.OnConn(&connHandler{
		listener:   listener,
		clientPool: clientPool,
		userPool:   userPool,
	})

	listener.Listen()
	select {}
}

type connHandler struct {
	listener transport.Listener

	clientPool *utils.Pool
	userPool   *utils.Pool
}

func (h *connHandler) ServeConn(conn transport.Conn) {
	client := common.NewClient(conn, h.clientPool)

	users.NewService(client, h.clientPool, h.userPool)
	messages.NewService(client, h.clientPool)
}
