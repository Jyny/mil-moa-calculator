package main

import (
	"log"
	"syscall/js"
)

var (
	window, document, body js.Value
)

func init() {
	window = js.Global()
	document = window.Get("document")
	body = document.Get("body")
}

func main() {
	log.Println("hello world")

	select {}
}
