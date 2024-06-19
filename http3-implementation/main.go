package main

import (
	"basic-http/udp"
	"fmt"
)

func main() {
	go func() {
		udp.Run()
	}()

	err := udp.Connect()
	fmt.Println("Connect completed")
	if err != nil {
		panic(err)
	}
}