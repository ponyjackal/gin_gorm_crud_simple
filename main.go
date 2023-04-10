package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Person struct {
	Id uint `json:"id"`
	FirstName string `jsong:"firstname"`
	LastName string `jsong:"lastname"`
}

func main() {
	db, _ := gorm.Open("sqlite3", "./gorm.db")
	defer db.Close()

	db.AutoMigrate(&Person{})

	p1 := Person{FirstName: "John", LastName: "Doe"}
	p2 := Person{FirstName: "Jane", LastName: "Smith"}

	db.Create(&p1)

	var p3 Person
	db.First(&p3)

	fmt.Println(p1.FirstName)
	fmt.Println(p2.LastName)
	fmt.Println(p3.LastName)
}