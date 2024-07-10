package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	CPF    string
	Email  string
	Age    int
	Active bool
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})

	return db
}

func AddStudent() {
	db := Init()

	student := Student{
		Name:   "Bendo",
		CPF:    "999.999.999-99",
		Email:  "example@example.com",
		Age:    21,
		Active: true,
	}

	if res := db.Create(&student); res.Error != nil {
		fmt.Println("Error to create student")
	}

	fmt.Println("Create student")
}
