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

	// data, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println(data)
	// fmt.Println(file)

	byteSlices := make([]byte, 2)
	numBytesRead, err := io.ReadFull(file, byteSlices)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(numBytesRead)
	fmt.Println(byteSlices)

}
