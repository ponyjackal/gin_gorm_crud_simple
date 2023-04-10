package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error
type Person struct {
	Id uint `json:"id"`
	FirstName string `jsong:"firstname"`
	LastName string `jsong:"lastname"`
}

func main() {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Person{})

	var people []Person

	db.Find(&people)

	r := gin.Default()
	r.GET("/", GetProjects)

	r.Run(":8080")
}

func GetProjects(c *gin.Context) {
	var people []Person

	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, people)
	}
}