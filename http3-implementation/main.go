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
	fmt.Println("udp connect completed")
	if err != nil {
		panic(err)
	}
}