package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	CPF    string `json:"cpf"`
	Email  string `json:"email"`
	Age    int    `json:"age"`
	Active bool   `json:"registration"`
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Student{})
	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student Student) error {
	if res := s.DB.Create(&student); res.Error != nil {
		return res.Error
	}
	fmt.Println("Create student")
	return nil
}

func (s *StudentHandler) GetStudent() ([]Student, error) {
	students := []Student{}
	err := s.DB.Find(&students).Error
	return students, err
}
