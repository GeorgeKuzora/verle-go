package main

import (
	"verle_go/pkg/handlers"
	"verle_go/pkg/server"
)

func main() {
	handlers.InitClients()
	handlers.RegisterHandlers()
	server.StartServer()
}
