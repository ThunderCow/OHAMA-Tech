package main

import (

	"os/signal"
	"os"
	"fmt"
	"time"
	"syscall"
)



	func main(){

		//cancel := context.WithCancel(context.Background())

		//defer cancel()

		quit:= make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT)
		//signal.Notify(sigs,os.Interrupt)
    
		go func(){
			<-quit
      
			//cancel()

			fmt.Printf("Du pressed kntrl + c. Bruker avbrytett eviglokke.")
			os.Exit(1)
		
    }()

		for i:= 0; ; i++{

			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}

	
/*
Programmet bruker 6.6 MG ram og nesten ingen CPU. Dette så vi vet å kjøre programmet i cmd,
og deretter sjekket vi hvor mye ram og cpu cmd brukte når programmet var under eksekvering. 
Når programmet var ikke i eksekvering bruktte cmd kun 6.6 MG RAM. Under ekskvering av pro-
grammet cmd 13.2 RAM og ikke noe CPU.
*/

