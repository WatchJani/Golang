package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("d1/file.txt")

	if err != nil {
		log.Println(err)
	}

	byteSlice, minBytes := make([]byte, 32), 1
	numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)

	if err != nil {
		log.Println(numBytesRead)
	}

	fmt.Printf("%s\n", byteSlice)

}
