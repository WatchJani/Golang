package main

import (
	"log"
	"os"
)

func main() {
	const originalPath string = "dir0"
	const newPath string = "dir4"

	_, err := os.Stat(originalPath)

	if os.IsNotExist(err) {
		err = os.Mkdir(originalPath, 0777)

		err = os.Rename(originalPath, newPath)

		if err != nil {
			log.Println(err)
		}
	}
}
