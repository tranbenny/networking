package http

import (
	"fmt"
	"net"
)

type Client struct {
}

type Response struct {
	Status int
	Body   []byte
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SendRequest(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return fmt.Errorf("error connecting to TCP server. %w", err)
	}
	defer conn.Close()

	// send a request
	req := []byte("hello world")
	if _, err := conn.Write(req); err != nil {
		return fmt.Errorf("error sending msg to server. %w", err)
	}

	return nil
}
