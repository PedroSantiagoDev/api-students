package main

import (
	"net/http"

	"github.com/PedroSantiagoDev/api-students/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/students", getStudents)
	e.POST("/students", createStudents)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func getStudents(c echo.Context) error {
	students, err := db.GetStudent()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get student")
	}
	return c.JSON(http.StatusOK, students)
}

func createStudents(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}

	if err := db.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create students")
	}
	return c.String(http.StatusCreated, "Create students")
}
