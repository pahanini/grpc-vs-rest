package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/user/:id", func(c echo.Context) error {
		return c.JSON(http.StatusOK,
			struct {
				ID    string
				Name  string
				Email string
			}{
				c.Param("id"),
				"foo",
				"foo@localhost",
			})
	})
	e.Logger.Fatal(e.Start(":8081"))
}
