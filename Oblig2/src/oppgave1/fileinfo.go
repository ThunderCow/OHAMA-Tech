
package main

import (
	"fmt"
	"log"
	"os"

)

var (
	fileInfo os.FileInfo
	fileMode os.FileMode
	err      error
)

type FileMode uint32


func fileIsRegular (){

	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		// check does not exist
    
		fmt.Println("The file does not exist")
	}


	if _, err := os.Stat("test.txt"); err == nil {
		// check file does exist
    
		fmt.Println("This is Regular file")
	}

}

func ReadLink(){


	if _, err := os.Lstat("test.txt"); err == nil {
    
    
		fmt.Println("The is a symbolic link")
    
	} else {
    
		fmt.Println("This is not a symbolic link")
		}

}


func sizeInBKM(){


file, err := os.Open("test.text")

if err != nil{
	return
	}

	defer file.Close()


	stat, err := file.Stat()

	if err != nil{
		return
	}



	var bytes = stat.Size()

	var kilobytes  = ((bytes)/1024)

	var megabytes = (float64)(kilobytes/1024)



	fmt.Printf("Size in bytes:", bytes)
	fmt.Printf("\n")
	fmt.Printf("Size in kilobytes:", kilobytes)
	fmt.Printf("\n")
	fmt.Printf("Size in megabytes:", megabytes)
	fmt.Printf("\n")

}


func main() {


	fileInfo, err = os.Stat("test.txt")
	if err != nil {
		log.Fatal(err)
	}


	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
  
  //close file
	file.Close()
	
  // OpenFile with more options. Last param is the permission mode
	// Second param is the attributes when opening

	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()


	//skriver ut info
	fmt.Println("Information about:", fileInfo.Name())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println( "Is not append only",)


	//methodcall
	fileIsRegular()
	ReadLink()
  sizeInBKM()
}





