package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)



func main() {
	fmt.Println("Concurrency in Go")

	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	
	for {
		_, err := io.WriteString(c, time.Now().Format("15:03:04\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}