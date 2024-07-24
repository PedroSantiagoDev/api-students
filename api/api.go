package api

import (
	"net/http"

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
	api.Echo.POST("/students", api.createStudents)
	//api.Echo.GET("/students/:id", api.getStudent)
	// api.Echo.PUT("/students/:id", api.updateStudents)
	// api.Echo.DELETE("/students/:id", api.deleteStudents)
}

func (api *API) Start() error {
	return api.Echo.Start(":8080")
}

func (api *API) getStudents(c echo.Context) error {
	students, err := api.DB.GetStudent()
	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get student")
	}
	return c.JSON(http.StatusOK, students)
}

func (api *API) createStudents(c echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}

	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create students")
	}
	return c.String(http.StatusCreated, "Create students")
}
