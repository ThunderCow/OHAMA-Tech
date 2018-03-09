/*The program is done which takes inputs from terminal & info manipulates further through channels.
*/
package main

import (
	"fmt"
)

//function that adds values of two inputs and send it to other function.
func addValue(input1 int, input2 int, chn2 chan<- int) {
	//writing to channel
	chn2 <- input1 + input2
}

//
//function to write user input to channels
func getUserInput(input1 int, input2 int, chn1 chan int,chn2 <-chan int,  chn3 chan int) {
	chn1 <- input1
	chn1 <- input2
	//receiving info from other func and read into this func.
	val := <-chn2
	chn3 <- val
}

//Main function
func main() {
	var input1 int 	//variable declaration
	var input2 int
	//channels declaration
	chn1 := make(chan int, 3)
	chn2 := make(chan int)
	chn3 := make(chan int)

	//taking inputs from terminal
	fmt.Scanln(&input1)
	fmt.Scan(&input2)

	//calling go functions
	go getUserInput(input1, input2, chn1, chn2, chn3)
	go addValue(input1, input2, chn2)
	go oppTreSIGINT()


	//print out the values on the terminal
	fmt.Println("Reading first input: ", <-chn1)
	fmt.Println("Reading second input: ", <-chn1)
	fmt.Println("Giving resulted value: ", <-chn3)
}

func oppTreSIGINT()  {
	//implementering av håndtering av SIGINT
	c := make(chan os.Signal, 0x2)
	signal.Notify(c, syscall.SIGINT)

	go func(){
		<- c
		fmt.Printf("program motatt SIGINT")
		os.Exit(0)
	}()

}
