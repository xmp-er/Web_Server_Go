package main

import (
	"fmt"
	"log"
	"net"
)

func perform_action(conn net.Conn) {
	inp := make([]byte, 1024)
	fmt.Println("Connection established, starting reading and writing operations")
	_, err := conn.Read(inp)
	if err != nil {
		fmt.Println("Error while reading data from client ", err)
	}
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nConnection established and response recieved\r\n"))
	conn.Close()
}

func main() {
	listener, err := net.Listen("tcp", ":1743")

	if err != nil {
		log.Fatal("Error in creating tcp listener", err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatal("Error in listener accepting connections", err)
		}

		fmt.Println("Server started and listening for data")
		go perform_action(conn)
	}

}
