package main

import (
	"fmt"
	"log"
	"os"
)

var (
	fileInfo os.FileInfo
	err      error
)

//Checks if the file is append or not
func checkFileIsAppend()  {
	filen := os.Args[1]
	_, err := os.Open(filen)
	switch  os.ModeAppend {
	case os.ModeAppend:
	if err == nil {
		fmt.Println("File is appended")
	} else {
		fmt.Println("File is not appended")
	}
	}
}

//Checks if the file is a device file or not
func checkFileDevice()  {
	filen := os.Args[1]
	_, err := os.Open(filen)
	switch  os.ModeDevice {
	case os.ModeDevice:
		if err == nil {
			fmt.Println("Is a device file")
		} else {
			fmt.Println("Is not a device file")
		}
	}
}

//Checks if the file is regular or not
func fileIsRegular (){
	filen := os.Args[1]
	if _, err := os.Stat(filen); os.IsNotExist(err) {
		// check does not exist
		fmt.Println("The file does not exist")
	}
	if _, err := os.Stat("test.txt"); err == nil {
		// check file does exist
		fmt.Println("This is Regular file")
	}

}

//Checks if the file is an symbolic link or not
func SymLink(){
	filen := os.Args[1]
	if _, err := os.Lstat(filen); err == nil {

		fmt.Println("The is a symbolic link")

	} else {

		fmt.Println("This is not a symbolic link")
	}

}

//Converts the file size into other units from bytes
func sizeInBKM(){
	filen := os.Args[1]
	file, err := os.Open(filen)

	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil{
		log.Fatal(err)
	}
	//bytes conversion into others with converting int to float64
	bytes := stat.Size()
	var b = float64(bytes)

	var kilobytes float64
	kilobytes = b / 1024

	var megabytes float64
	megabytes = kilobytes / 1024

	var gigabytes float64
	gigabytes = megabytes /1024

	//print out the following conversions
	fmt.Println("File size in bytes ", bytes, "bytes")
	fmt.Println("File size in kilobytes ", kilobytes, "KB")
	fmt.Println("File size in megabytes ", megabytes, "MB")
	fmt.Println("File size in gigabytes", gigabytes, "GB")

}

//main function
func main() {

	//filen gets the fil wanted to check as a argument in terminal
	filen := os.Args[1]

	//file_info function to open and read file
	fileInfo, err = os.Stat(filen)
	if err != nil {
		log.Fatal(err)
	}

	//print out info and calling methods
	fmt.Println("Information about:", fileInfo.Name())
	sizeInBKM()
	fmt.Println("Is Directory: ", fileInfo.IsDir())
	fileIsRegular()
	fmt.Println("Permissions:", fileInfo.Mode())
	checkFileIsAppend()
	checkFileDevice()
	SymLink()
}
