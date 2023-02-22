package main

import (
	ormscripts "GORM-Database-Example/ORM-Scripts"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	fmt.Println("This is a GORM Example Project.")
	ormscripts.Orgnize()

}
