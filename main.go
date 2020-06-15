package main

import (
	"github.com/undefined7887/telepuz-backend/base/log"
	"github.com/undefined7887/telepuz-backend/base/network"
	"os"
	"regexp"
)

var addrRegexp = regexp.MustCompile("")

func main() {
	logger := log.NewDefaultLogger("")

	addr := os.Args[1]
	if !addrRegexp.MatchString(addr) {
		logger.Fatal("Wrong address format: address should be like ip:port")
	}

	listener := network.NewWebsocketListener(logger, "/", addr)

}
