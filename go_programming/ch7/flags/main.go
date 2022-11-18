package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 7*time.Second, "sleep period")

func main() {
	fmt.Println("Flags in Go")
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
