package main

import (
	"log"
	"os"
)

func main() {

	err := os.Mkdir("dir0", 0777)

	if err != nil {
		log.Println(err)
	}

	err = os.MkdirAll("div123/derer/eqasdasd/", 0777)
	if err != nil {
		log.Println(err)
	}

}
