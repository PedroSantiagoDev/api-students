package api

import (
	"github.com/PedroSantiagoDev/api-students/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type API struct {
	Echo *echo.Echo
	DB   *db.StudentHandler
}

func NewServer() *API {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	database := db.Init()
	studentDB := db.NewStudentHandler(database)

	return &API{
		Echo: e,
		DB:   studentDB,
	}
}

func (api *API) ConfigureRoutes() {
	api.Echo.GET("/students", api.getStudents)
	api.Echo.POST("/students", api.createStudent)
	api.Echo.GET("/student/:id", api.getStudent)
	// api.Echo.PUT("/students/:id", api.updateStudents)
	// api.Echo.DELETE("/students/:id", api.deleteStudents)
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}
