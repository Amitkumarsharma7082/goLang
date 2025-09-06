package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		//! create a file
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("error occured")
		}
		defer file.Close()

		//! write content
		content := "hello go lang"
		_, error := io.WriteString(file, content) // byte in file

		if error != nil {
			fmt.Println("error occured in adding content")
			return
		}
		fmt.Println("created")
	*/

	//! read a file
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("error when file open", err)
		return
	}
	defer file.Close()
	/*
		//! buffer use
		// buffer is use for temp storage hold the data for sometime

		buffer := make([]byte, 1024) // take a byte in reading file (using Slice)
		// read till
		for {
			n, err := file.Read(buffer) //
			if err == io.EOF {
				fmt.Println("end of file")
				break
			}
			if err != nil {
				fmt.Println("error when reading", err)
				return
			}
			// process the reading file till 0:n(0 to n)
			fmt.Println(string(buffer[:n]))
		}
	*/
	// ioutil.ReadFile
	// after evolve the go we don't use ioutil.ReadFile

	//os.ReadFile
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("reading file ioutilReadfil", err)
		return
	}
	fmt.Println(string(content))

}
