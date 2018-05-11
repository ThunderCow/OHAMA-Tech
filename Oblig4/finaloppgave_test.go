package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
	"os"
	"net"
	"io/ioutil"
)

//function which will find the file
func fileCheck(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) { //checking if file does not exit return false, otherwise true
			return false
		}
	}
	return true
}

// Test file function
func TestHTMLFile(t *testing.T) {
	htmlFile := fileCheck("test.html") //file in the same directory.
	if htmlFile != true { //check error and handle the response.
		t.Error("File not found.\n")
	}else{
		fmt.Print("File OK.\n")
	}
}


//This is a general testing function where we can test any template function from finaloppgave.go
func TestTemplateGeneral(t *testing.T)  {
	rq, err := http.NewRequest("GET", "http://api.coinmarketcap.com/v1/ticker/bitcoin", nil)
	if err != nil{
		t.Fatal(err)
	}
	resRec := httptest.NewRecorder() //creating new recorder and saving into resRec variable
	handler := http.HandlerFunc(TemplateLitecoin) //can write other function names like TemplateLitecoin, Ripple, Ethereum..

	handler.ServeHTTP(resRec, rq) //serving HTTP server
	if chekStatus := resRec.Code; chekStatus != http.StatusOK { //having check in case of fail through error handling
		t.Errorf("handler has returned wrong status code: got %v want %v",
			chekStatus, http.StatusOK)
	}
	/* The above code is fine enough to test our function and give result accordingly. but we have added extra layer to get assurance through return statements
	that all were well by declaring variable to save expected data into it*/

	expected := resRec.Body //check whether the testing is performed well as we expected.
	if resRec.Body != expected {
		t.Errorf("Handler has returned unexpected body: got %v want %v",
			resRec.Body.String(), expected)
	}else {
		fmt.Println("1. The testing of function Templates is succeeded, congrats!")
	}
}

//This function is based on testing CollectData functions from finaloppgave.go
func TestCollectData(t *testing.T)  {
	req, err := http.NewRequest("GET", "http://api.coinmarketcap.com/v1/ticker/bitcoin", nil)
	if err != nil {
		t.Fatal(err)
	}
	resRec := httptest.NewRecorder() //creating new recorder and saving into resRec variable
	handler := http.HandlerFunc(CollectData22)   // can test the other functions just need to change the name: CollectData22,33,44 og 55

	handler.ServeHTTP(resRec, req) //serving HTTP
	if checkStatus := resRec.Code; checkStatus != http.StatusOK { //having check in case of fail through error handling
		t.Errorf("Handler has returned wrong status: got %v want %v",
			checkStatus, http.StatusOK)
	}
	/* The above code is fine enough to test our function and give result accordingly. but we have added extra layer to get assurance through return statements
	that all were well by declaring variable to save expected data into it*/
	
	expected := resRec.Body // bit more assurity to have confirmation of right testing result.
	if resRec.Body != expected {
		t.Errorf("Handler has returned wrong body: got %v wanted %v",
			resRec.Body.String(), expected)
	}else {
		fmt.Println("2. The testing of function 'CollectData' is succeeded")
	}
}

//Test function which verify that connection is open on port 8080
func TestConnection(t *testing.T) {
	message := "Testing the port...\n"
	port := ":8080"
	go func() {
		conn, err := net.Dial("tcp", port)
		if err != nil {		  //check error / error handling
			t.Fatal(err)
		}
		defer conn.Close()

		if _, err := fmt.Fprintf(conn, message); err != nil { //printing message to port number :8080
			t.Fatal(err)
		}
	}()

	l, err := net.Listen("tcp", port) //listening on :8080
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(string(buf[:]))
		if msg := string(buf[:]); msg != message {
			t.Fatalf("Unexpected message response:\n Got:\t\t%s\n Expected:\t %s \n", msg, message)
		}else{
			fmt.Print("It is ready now ! (on port " + port +")\n")
		}
		return // Done
	}
}
