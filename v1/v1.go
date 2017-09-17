package v1

import (
	"net/http"

	"github.com/aokabin/tabibayashi/middleware"
	"github.com/labstack/echo"
)

func EchoHandler() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")
	v1.Use(middleware.V1GroupFunc)

	useV1Group(v1)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}
