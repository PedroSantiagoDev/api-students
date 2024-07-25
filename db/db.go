package db

import (
	"github.com/rs/zerolog/log"
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
		log.Fatal().Err(err).Msgf("Failed to initialize SQlite %s", err.Error())
	}

	db.AutoMigrate(&Student{})
	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student Student) error {
	if res := s.DB.Create(&student); res.Error != nil {
		log.Error().Msg("Failed to create Student")
		return res.Error
	}
	log.Info().Msg("Create Student")
	return nil
}

func (s *StudentHandler) GetStudents() ([]Student, error) {
	students := []Student{}
	err := s.DB.Find(&students).Error
	return students, err
}

func (s *StudentHandler) GetStudent(id int) (Student, error) {
	student := Student{}
	err := s.DB.First(&student, id)
	return student, err.Error
}

func (s *StudentHandler) UpdateStudent(updatingStudent Student) error {
	return s.DB.Save(&updatingStudent).Error
}
