package db

import (
	"github.com/PedroSantiagoDev/api-students/schemas"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StudentHandler struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to initialize SQlite %s", err.Error())
	}

	db.AutoMigrate(&schemas.Student{})
	return db
}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student schemas.Student) error {
	if res := s.DB.Create(&student); res.Error != nil {
		log.Error().Msg("Failed to create Student")
		return res.Error
	}
	log.Info().Msg("Create Student")
	return nil
}

func (s *StudentHandler) GetStudents() ([]schemas.Student, error) {
	students := []schemas.Student{}
	err := s.DB.Find(&students).Error
	return students, err
}

func (s *StudentHandler) GetStudent(id int) (schemas.Student, error) {
	student := schemas.Student{}
	err := s.DB.First(&student, id)
	return student, err.Error
}

func (s *StudentHandler) GetFilteredStudent(active bool) ([]schemas.Student, error) {
	filteredStudents := []schemas.Student{}
	err := s.DB.Where("active = ?", active).Find(&filteredStudents)
	return filteredStudents, err.Error
}

func (s *StudentHandler) UpdateStudent(updatingStudent schemas.Student) error {
	return s.DB.Save(&updatingStudent).Error
}

func (s *StudentHandler) DeleteStudent(student schemas.Student) error {
	return s.DB.Delete(&student).Error
}
