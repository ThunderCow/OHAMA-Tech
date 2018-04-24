
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"bufio"
)

//Function to check any error.
func CheckError(err error){
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}
}

// function TCP to launch server and listen on port 17.

func serverTCP()  {
	// Listen for incoming connections.
	fmt.Println("Launching TCP Server...")
	l, err := net.Listen("tcp", ":17")
	if err != nil {
		fmt.Printf("Error listening: %v", err.Error())
		return
	}

	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + ":17")
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
		//go handler(conn)
	}
}

// Function UDP to listen on the same port and send response.

func serverUDP()  {

	fmt.Println("Launching UDP Server...")
	/* prepare a address at any address at port 17*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":17")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		newmessage := strings.ToUpper(string(buf[0:n]))
		// send new string back to client
		ServerConn.WriteToUDP([]byte(newmessage + "\n"),addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

//to handle the connection request which we get.
func handleRequest(conn net.Conn)  {
	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		//newmessage := strings.ToUpper(message)
		// send new string back to client
		conn.Write([]byte("hei på degasda asda d dasa s ;)" + "\n"))
	}
	conn.Close()
}

func main()  {

//declaring variable to take input from terminal.
var input1 int
//To ask which server is meant to be
fmt.Println("which server are you meant to be: ")

fmt.Println("Press 1 for TCP or Press 2 for UDP: ")
//taking input
fmt.Scanln(&input1)
//checking whether input is correct if yes then run desired server.
if input1==1 {
	serverTCP()
}
if input1 == 2 {
	serverUDP()
}else {  // sending message to select right choice.
	fmt.Println("You must choose 1 or 2 to select right server.")
}

}
