package main

import (
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/network"
	"github.com/undefined7887/telepuz-backend/rand"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
	"github.com/undefined7887/telepuz-backend/service/format"
	"github.com/undefined7887/telepuz-backend/service/messages"
	"github.com/undefined7887/telepuz-backend/service/users"
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

	clientPool := repository.NewPool()
	userPool := repository.NewPool()

	listener := network.NewWebsocketListener(logger, "/", addr)
	listener.Handle(&connHandler{
		listener:   listener,
		clientPool: clientPool,
		userPool:   userPool,
	})

	select {}
}

type connHandler struct {
	listener network.Listener

	clientPool *repository.Pool
	userPool   *repository.Pool
}

func (h *connHandler) ServeConn(conn network.Conn) {
	client := &service.Client{
		Id:         rand.HexString(format.IdLength),
		ClientPool: h.clientPool,
		Conn:       conn,
	}
	h.clientPool.Add(client.Id, client)

	users.NewService(client, h.clientPool, h.userPool)
	messages.NewService(client, h.clientPool, h.userPool)
}
