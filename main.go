package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}

	defer listener.Close()
	fmt.Println(
		"Listener accepting connections on localhost 8080")

	for {
		conn, err := listener.Accept()
		if err == nil {
			go proxyConnection(conn)
		}
	}
}

func writeForward(
	incomingConnection net.Conn,
	outgoingConnection net.Conn) {
	defer outgoingConnection.Close()
	for {
		buffer := make([]byte, 5000)
		dataSize, err := incomingConnection.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		data := buffer[:dataSize]
		_, err = outgoingConnection.Write(data)
		if err != nil {
			panic(err)
		}
	}
}

func writeBackward(
	outgoingConnection net.Conn,
	incomingConnection net.Conn) {
	defer incomingConnection.Close()
	for {
		buffer := make([]byte, 5000)
		dataSize, err := outgoingConnection.Read(buffer)
		if err != nil {
			panic(err)
		}
		data := buffer[:dataSize]
		_, err = incomingConnection.Write(data)
		if err != nil {
			panic(err)
		}
	}
}

func proxyConnection(inConn net.Conn) {
	outConn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		panic(err)
	}
	go writeForward(inConn, outConn)
	go writeBackward(outConn, inConn)
}
