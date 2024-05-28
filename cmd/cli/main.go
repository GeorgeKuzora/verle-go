package main

import (
	"verle_go/pkg/handlers"
	"verle_go/pkg/log"
	"verle_go/pkg/server"
)

func main() {
	log.LogInit()
	handlers.InitClients()
	handlers.RegisterHandlers()
	server.StartServer()
}
