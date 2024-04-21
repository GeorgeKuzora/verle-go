package main

import (
	"verle_go/pkg/handlers"
	"verle_go/pkg/server"
	"verle_go/pkg/sheets"
	"verle_go/pkg/weeek"
)

func main() {
	weeek.InitClient()
	sheets.InitClient()
	handlers.RegisterHandlers()
	server.StartServer()
}
