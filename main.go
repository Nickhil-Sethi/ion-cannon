package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen(
		"tcp", "localhost:8080")

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	fmt.Println(
		"Listener accepting connections on localhost 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(
				"Error accepting from connection!")
			continue
		}
		go proxyConnection(conn)
	}
}

func writeForward(
	incomingConnection net.Conn,
	outgoingConnection net.Conn) {

	for {
		io.Copy(outgoingConnection, incomingConnection)
	}
}

func writeBackward(
	outgoingConnection net.Conn,
	incomingConnection net.Conn) {

	for {
		io.Copy(incomingConnection, outgoingConnection)
	}
}

func proxyConnection(inConn net.Conn) {
	outConn, err := net.Dial(
		"tcp", "google.com:80")

	if err != nil {
		panic(err)
	}

	go writeForward(inConn, outConn)
	go writeBackward(outConn, inConn)
}
