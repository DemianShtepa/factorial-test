package main

import (
	"github.com/DemianShtepa/factorial-test/internal/http/server"
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	srv := server.InitializeServer()

	return http.ListenAndServe("127.0.0.1:8989", srv)
}
