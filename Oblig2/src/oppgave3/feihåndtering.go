package main

import (
	"fmt"
)
//to handle an error
func check (e error){
	if e != nil{
		panic(e)
	}
}
func main() {
f, err := os.OpenFile(filename, os.O_WRONLY, 0777)
	//check if error exist to call check error
  if err != nil {
		check(err)
	}
}
