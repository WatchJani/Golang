package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("d1/test2.txt")
	if err != nil {
		log.Println(err)
	}

	defer file.Close()

	counter := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		counter++
		err := scanner.Err()
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("Line %v > Read %d char > %v\n", counter, len(scanner.Text()), string(scanner.Text()))
	}

}
