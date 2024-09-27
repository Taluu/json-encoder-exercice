package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	domain "github.com/Taluu/json-encoder-exercise/pkg"
	"github.com/Taluu/json-encoder-exercise/pkg/di"
	_ "github.com/Taluu/json-encoder-exercise/pkg/encoder/base64"
	_ "github.com/Taluu/json-encoder-exercise/pkg/http"
	"github.com/Taluu/json-encoder-exercise/pkg/signature/hmac"
)

func main() {
	host := flag.String("host", "localhost", "Set the host")
	port := flag.Uint("port", 8080, "The port to listen to")
	key := flag.String("key", "", "key to use to compute signatures")

	flag.Parse()

	if *key == "" {
		// stop everything, there's no point for this if we do not have a key
		// A possible amelioration would be to be able to setup a proper env with
		// proper env variable, such as using the viper lib, but this is out of
		// scope for this poc.
		panic("empty key")
	}

	// have to declare the service here, because of the string param
	di.Provide[domain.SignatureHandler](func() domain.SignatureHandler {
		return hmac.NewSignatureHandler(*key)
	})

	addr := fmt.Sprintf("%s:%d", *host, *port)

	di.RegisterHTTP("/encrypt")
	di.RegisterHTTP("/decrypt")

	di.RegisterHTTP("/sign")

	log.Printf("Starting to listen on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
