package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr = flag.String("addr", "0.0.0.0:8080", "address:port to listen on")
	root = flag.String("root", ".", "path for serving static files")
)

func init() {
	flag.Parse()
}

func main() {
	http.Handle(
		"/",
		http.FileServer(
			http.Dir(*root),
		),
	)

	log.Printf("Server running at http://%s/\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
