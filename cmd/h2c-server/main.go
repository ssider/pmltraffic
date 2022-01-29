package main

import (
	"pml/traffic-test/pkg/body"
	"flag"
	"fmt"
	"golang.org/x/net/http2"
	"log"
	"net"
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

	lsner, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	log.Printf("listening on %d\n", *port)
	if err != nil {
		log.Fatal(err)
	}

	srv := &http2.Server{}
	for {
		conn, err := lsner.Accept()
		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		}
		go func() {
			opts := &http2.ServeConnOpts{Handler: handler}
			srv.ServeConn(conn, opts)
		}()
	}
}
