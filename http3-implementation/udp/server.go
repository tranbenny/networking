package udp

import (
	"fmt"
	"net"
	"os"
)

const (
	ServerAddress = "localhost:7000"
	bufferSize = 65507
)


func Run() {
	addr, err := net.ResolveUDPAddr("udp", ServerAddress)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	for {
		var buf [bufferSize]byte
		_, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}

		os.Stdout.Write([]byte("recevied request\n"))
		os.Stdout.Write(buf[0:])
		os.Stdout.Write([]byte("\n"))

		resp := "ack\n"
		if _, err = conn.WriteToUDP([]byte(resp), addr); err != nil {
			os.Stderr.Write([]byte(err.Error()))
		}
	}
}
