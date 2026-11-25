package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// 1, Config: read from environment variables, with defaults
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	name := os.Getenv("NAME")
	if name == "" {
		name = "server-1"
	}

	// 2. Handler: runs once for every incoming request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s listening on port %s", name, port)
		fmt.Fprintf(w, "hello from %s\n", name)
	})

	log.Printf("Starting server %s on port %s", name, port)
	log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
