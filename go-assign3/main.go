package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	filename := os.Args[1]
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}

	data := make([]byte, 100)     //read into byte string
	count, err := file.Read(data) //get length of file and check readable
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The file", filename, "has", count, "and its contents are:", string(data))

}
