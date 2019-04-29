package main

import (
	"github.com/ericzhang-cn/go-ps/server"
)

func main() {
	c := server.Config{
		IPAddress: "0.0.0.0",
		Port:      7777,
		BadgerDir: "./data",
	}
	s := server.PsServer{
		Rank: 1,
		C:    &c,
	}
	s.Serve()
}
