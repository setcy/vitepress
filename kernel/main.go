package main

import (
	"github.com/setcy/wiki/kernel/model"
	"github.com/setcy/wiki/kernel/server"
	"github.com/setcy/wiki/kernel/sql"
)

func main() {
	model.InitConf()
	sql.InitDatabase()
	server.Serve()
}
