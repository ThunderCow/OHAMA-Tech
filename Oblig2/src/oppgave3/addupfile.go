
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//to handle an error
func check (e error){
	if e != nil{
		panic(e)
	}
}

// var. declaration
var a string
var b string

//Main function
func main() {
	

	//implementering av håndtering av SIGINT
	c := make(chan os.Signal, 0x2)
	signal.Notify(c, syscall.SIGINT)

	go func(){
		<- c
		fmt.Printf("program motatt SIGINT")
		os.Exit(0)
	}()

	//readValue from terminal
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Println("Reading number", a, b, "\n")	//output the reading

	//open a file to read it and check for error too
	f, err := os.Create("test1.txt")
	check(err)
	defer f.Close()

	//bytes conversion & write on a file
	bs := []byte(a + "\n" + b)
	n1, err := f.Write(bs)
	check(err)
	fmt.Printf("wrote %d bytes \n", n1)

	//syncing to ensure all is done
	f.Sync()

}
