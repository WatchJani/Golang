package database

import (
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Println(e)
	}

}

func Exist(name string) bool {
	_, err := os.Stat(name)

	return os.IsNotExist(err)
}
