package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-messages:
			// broadcast all incoming message to all
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}

	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	CheckError(err, "")

	go broadcaster()
	for {
		conn, err := listener.Accept()
		CheckError(err, "skip")
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)

	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// ignoring potential errors
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // ignore network issues
	}
}

func CheckError(e error, optional string) {
	if optional != "" {
		if e != nil {
			log.Fatal(e)
			return
		}
	} else {
		if e != nil {
			log.Fatal(e)
		}
	}
}
