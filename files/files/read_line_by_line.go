package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("d1/test2.txt")

	if err != nil {
		log.Println(err)
	}

	reader := bufio.NewReader(file)
	line, counter := "", 0

	for {
		line, err = reader.ReadString('\n')
		if err != nil{
			break
		}
		counter++
		fmt.Printf("Line %v > Read %d character > %v", counter, len(line), string(line))
	}

	if err != io.EOF{
		fmt.Println(" > Filed!: %v\n", err)
	}

}
