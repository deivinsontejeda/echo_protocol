package main

import (
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()
	for {
		// Wait for a connection.
		log.Print("Waiting connection")
		conn, err := l.Accept()

		if err != nil {
			log.Fatal(err)
		}

		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			buffer := make([]byte, 1024)
			bytesRead, _ := conn.Read(buffer)
			log.Print(string(buffer))
			log.Print(bytesRead)
			response := "Hi there, current time"
			c.Write([]byte(response))
			c.Close()
		}(conn)

	}
}
