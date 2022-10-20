package main

import (
	"log"
	"os"
)

func main() {
	FileInfo, err := os.Create("text.txt")

	if err != nil {
		log.Println(err)
	}

	FileInfo.Close()

	log.Println(FileInfo.Name())

}
