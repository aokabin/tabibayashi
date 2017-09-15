package v1

import (
	"net/http"

	"github.com/aokabin/tabibayashi/middleware"
	"github.com/labstack/echo"
)

func EchoHandler() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")
	v1.Use(middleware.SampleFunc)

	useSampleGroup(v1)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/visit", visit)
	// e.Logger.Fatal(e.Start(":1323"))
	return e
}
