package http

import (
	"fmt"
	"io"
	"net"
)

type TCPServer struct {
	host string
	port int
}

func NewTCPServer(host string, port int) *TCPServer {
	return &TCPServer{
		host: host,
		port: port,
	}
}

func (t *TCPServer) Run() {
	fmt.Println("starting server...")
	addr := fmt.Sprintf("%s:%d", t.host, t.port)
	server, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("received error. crashing server")
			panic(err)
		}
		if err := handleConnection(conn); err != nil {
			panic(err)
		}

	}
}

func handleConnection(conn net.Conn) error {
	defer conn.Close()
	fmt.Println("received connection")
	resp, err := io.ReadAll(conn)
	if err != nil {
		return fmt.Errorf("error reading incoming request. %w", err)
	}
	fmt.Println(string(resp))
	return nil
}
