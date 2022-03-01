package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in golnag")
	content := "This needs to go in a file - Zeeshan"

	file, err := os.Create("/Users/zeesh/OneDrive/Desktop/mygolang/Golang Basics/18files/mygofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("/Users/zeesh/OneDrive/Desktop/mygolang/Golang Basics/18files/mygofile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}