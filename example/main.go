package main

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/jquiterio/echopongo"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Render = echopongo.Default()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", pongo2.Context{"name": "World"})
	})
	e.Logger.Fatal(e.Start(":8000"))
}
