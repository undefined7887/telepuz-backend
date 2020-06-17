package main

import (
	"github.com/undefined7887/telepuz-backend/log"
	"github.com/undefined7887/telepuz-backend/network/websocket"
	"github.com/undefined7887/telepuz-backend/repository"
	"github.com/undefined7887/telepuz-backend/service"
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

	listener := websocket.NewListener(logger, "/", addr)
	listener.Handle(service.NewConnHandler(listener, clientPool, userPool))

	select {}
}
