package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/Taluu/json-encoder-exercise/pkg/di"
	_ "github.com/Taluu/json-encoder-exercise/pkg/encoder/base64"
	_ "github.com/Taluu/json-encoder-exercise/pkg/http"
)

func main() {
	host := flag.String("host", "localhost", "Set the host")
	port := flag.Uint("port", 8080, "The port to listen to")

	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	di.RegisterHTTP("/encrypt")
	di.RegisterHTTP("/decrypt")

	log.Printf("Starting to listen on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
