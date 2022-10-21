package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Println(e)
	}
}

func main() {

	d1 := []byte("Hello\nGolang\n")
	err := ioutil.WriteFile("d1/file.txt", d1, 0644)
	check(err)

	f, err := os.Create("d1/test2.txt")
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Println(n2)

	d3, err := f.WriteString("writes\n")
	check(err)
	fmt.Println(d3)

	f.Sync()
}
