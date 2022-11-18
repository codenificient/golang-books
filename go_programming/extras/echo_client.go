package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server-address")
		os.Exit(1)
	}
	serverAddress := os.Args[1]
	conn, err := net.Dial("tcp", serverAddress+":10000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userin := bufio.NewScanner(os.Stdin)
	serverin := bufio.NewScanner(conn)
	for {
		fmt.Print("Enter Text: ")
		userin.Scan()
		s := userin.Text()
		if strings.ToLower(s) == "exit" {
			break
		} else {
			fmt.Fprintln(conn, s)
			fmt.Printf("Sent: %s\n", s)
			serverin.Scan()
			fmt.Printf("Received: %s\n", serverin.Text())
		}
	}
	fmt.Println("Client is exitting...")
}
