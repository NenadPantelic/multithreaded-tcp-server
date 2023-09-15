package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func do(conn net.Conn) {
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Processing the request")
	time.Sleep(5 * time.Second)

	// write the byte slice
	conn.Write([]byte("HTTP/1.1 200\r\n Hello, World!\r\n"))
	conn.Close()
}

func main() {
	// 1. a process reserves the port 1729
	listener, err := net.Listen("tcp", ":1729")
	if err != nil {
		log.Fatal(err)
	}

	for {
		log.Println("Waiting for a client to connect")
		// 2. wait for the connection
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Client connected")
		// 3. process the request (read/write/close)
		// do(connection)
		// 5. paralize the processing of the requests
		go do(connection)
		fmt.Println("Hello world!")
	}

}
