package database

import (
	"log"
	"os"
)

//***************************????????????

// type Table struct {
// 	Name string
// 	Row  map[string]interface{}
// }

// type DataBase struct {
// 	Name  string
// 	Table []Table
// }

// func New(name string) *DataBase {
// 	return &DataBase{
// 		Name: name,
// 	}
// }

func CreateDB(name string) {
	if Exist(name) {
		err := os.Mkdir(name, 0777)
		check(err)
	} else {
		log.Println("This name is already Exist!")
	}
}

func DeleteDB(name string) {
	if !Exist(name) {
		err := os.RemoveAll(name)
		check(err)
	} else {
		log.Println("This name of directorium is not exist")
	}
}

func Rename(originalPath, newPath string) {
	if !Exist(originalPath) {
		err := os.Rename(originalPath, newPath) //KAKO DA FUNKCIJI POSALJEM SAMO OVU LINIJU KODA
		check(err)
	} else {
		log.Println("This name of directorium is not exist")
	}
}
