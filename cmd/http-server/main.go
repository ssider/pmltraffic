package main

import (
	"pml/traffic-test/pkg/body"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	respEnvs = []string{
		"POD_NAME",
		"POD_NAMESPACE",
	}
)

func main() {
	port := flag.Int("port", 80, "http server listening port")
	flag.Parse()
	handler := http.HandlerFunc(body.HTTPHandler)

	h1s := &http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: handler,
	}
	log.Printf("listening on %d\n", *port)
	log.Fatal(h1s.ListenAndServe())
}
