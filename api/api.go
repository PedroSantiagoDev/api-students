package api

import (
	"net/http"

	"github.com/PedroSantiagoDev/api-students/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type API struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db := db.Init()

	return &API{
		Echo: e,
		DB:   db,
	}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", getStudents)
	api.Echo.POST("/students", createStudents)
	api.Echo.GET("/students/:id", getStudents)
	// api.Echo.PUT("/students/:id", updateStudents)
	// api.Echo.DELETE("/students/:id", deleteStudents)
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

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
