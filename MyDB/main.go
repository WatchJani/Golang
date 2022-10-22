package main

import (
	db "root/database"
)

func main() {
	db.CreateDB("Mongo")
	//db.DeleteDB("Mongo")
	//db.Rename("Mongo", "MySQL")
}
