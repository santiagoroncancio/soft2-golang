package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000/"},
	}))

	e.GET("/", LoadData)
	e.Start(":8080")
}

func LoadData(c echo.Context) error {
	a := []struct {
		name string `json:"name"`
	}{
		{"Andres"},
		{"Santiago"},
		{"Entrada"},
	}
	return c.JSON(http.StatusOK, a)
}
