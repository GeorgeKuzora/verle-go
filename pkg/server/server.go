package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	// Start the HTTP server.
	port := ":8080"
	fmt.Printf("Server is listening on port %s...\n", port)
	http.ListenAndServe(port, nil)
}
