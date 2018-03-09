package main

import (
	"fmt"
	"os"
)
//to handle an error
func check (e error){
	if e != nil{
		panic(e)
	}
}

//Main function
func main() {

	filename := os.Args[1]
	//readValue from terminal

	//open a file to read it and check for error too

	a := os.Args[2]
	b := os.Args[3]

	f, err := os.OpenFile(filename, os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	check(err)
	defer f.Close()

	sb := []byte(a + "\n" + b)
	strValue, err := f.Write(sb)
	fmt.Println("Written values on a file", strValue)

}
