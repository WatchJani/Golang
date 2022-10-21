package main

import (
	"fmt"
	"os"
)

func main() {
	
	_, err := os.Stat("dir0")

	if os.IsNotExist(err) {
		_ = os.Mkdir("dir0", 0777)
		// if err != nil {
		// 	log.Println(err)
		// }
		fmt.Println("Directory created!")
	} else {
		fmt.Println("Directory already exist!")
	}
}
