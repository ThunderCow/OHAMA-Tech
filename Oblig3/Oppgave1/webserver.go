package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayHelloClient (w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)  //print form information
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("Key: ", k)
		fmt.Println("Val: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w,"Hello, Client!") // send data to client
}

func main()  {
	http.HandleFunc("/", sayHelloClient) // set router
	err := http.ListenAndServe(":8080", nil) // set listen port 8080
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
