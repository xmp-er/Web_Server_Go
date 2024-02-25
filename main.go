package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func perform_action(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	inp := make([]byte, 1024)
	fmt.Println("Connection established, starting reading and writing operations")
	_, err := conn.Read(inp)
	if err != nil {
		fmt.Println("Error while reading data from client ", err)
	}
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nConnection established and response recieved\r\n"))
	conn.Close()
}

func validatePortFlag(s string) bool {
	if s[0] != ':' || len(s) != 5 {
		return false
	}
	for _, val := range s {
		if val != ':' && (val > '9' || val < '0') {
			return false
		}
	}
	return true
}

func validateConnMethod(s string) bool {
	if s != "tcp" && s != "tcp4" && s != "tcp6" && s != "unix" && s != "unixpacket" {
		return false
	}
	return true
}

func main() {

	sigShut := make(chan os.Signal, 1)
	var wg sync.WaitGroup

	var port string
	var connMethod string
	flag.StringVar(&port, "p", ":8000", "Port number to be used in format :<port-number>")
	flag.StringVar(&connMethod, "c", "tcp", "Conection method to be specified either tcp/tcp6/tcp4/unix/unixpacket")
	flag.Parse()

	if !validatePortFlag(port) {
		log.Fatal("Please enter the port number in correct format as :<port-number>")
	}

	if !validateConnMethod(connMethod) {
		log.Fatal("Please enter a valid method between tcp/tcp6/tcp4/unix/unixpacket")
	}

	listener, err := net.Listen(connMethod, port)

	fmt.Println("Listening at port", port, "and the method is", connMethod)

	if err != nil {
		log.Fatal("Error in creating listener", err)
	}

	signal.Notify(sigShut, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-sigShut:
				fmt.Println(1)
				sigShut <- os.Kill
				return
			default:
				conn, err := listener.Accept()
				if err != nil {
					log.Fatal("Error in listener accepting connections", err)
				}

				fmt.Println("Server started and listening for data")
				wg.Add(1)
				go perform_action(conn, &wg)
			}
		}
	}()

	<-sigShut

	fmt.Println("\n Closing the listener \n")
	listener.Close()

	wg.Wait()
}
