package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	staticFilesHandler()
	port := ":8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	http.ListenAndServe(port, nil)
}

func staticFilesHandler() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
