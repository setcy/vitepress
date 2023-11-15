package main

import (
	"github.com/setcy/wiki/kernel/model"
	"github.com/setcy/wiki/kernel/server"
)

func main() {
	model.InitConf()
	server.Serve()
}
