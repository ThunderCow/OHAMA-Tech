package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	conn, err := net.Dial("udp", "localhost:17")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}

	// read in input from stdin
	reader := bufio.NewReader(os.Stdin)

	
	conn.Close()
}
