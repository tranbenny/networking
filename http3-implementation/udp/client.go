package udp

import (
	"bufio"
	"fmt"
	"net"
)

func Connect() error {
	addr, err := net.ResolveUDPAddr("udp", ServerAddress)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	defer conn.Close()
	if err != nil {
		return err
	}

	if _, err = conn.Write([]byte("hello world")); err != nil {
		return err
	}

	resp, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Printf("response: %s\n", resp)
	return nil
}
