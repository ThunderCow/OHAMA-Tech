package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {

	// connect to this socket
	conn, _ := net.Dial("tcp", "localhost:17")
	// read in input from stdin
	reader := bufio.NewReader(os.Stdin)

	
}
