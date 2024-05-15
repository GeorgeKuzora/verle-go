package main

import (
	"verle_go/pkg/handlers"
	"verle_go/pkg/server"
	"verle_go/pkg/sheets"
)

func main() {
	sheets.InitClient()
	handlers.RegisterHandlers()
	server.StartServer()
}
