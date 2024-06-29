package main

import (
	"http2/http"
	"time"
)

func main() {
	s := http.NewTCPServer("localhost", 9000)
	go s.Run()

	c := http.NewClient()
	c.SendRequest("localhost:9000")

	// TODO: split out the main function and remove this sleep
	time.Sleep(time.Second * 5)
}
