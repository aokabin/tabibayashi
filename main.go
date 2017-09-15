package main

import (
	"context"

	"github.com/aokabin/tabibayashi/v1"
	"github.com/labstack/echo"
)

var (
	s *echo.Echo
)

func main() {
	defer shutdownServer()

	startServe()
	// e := v1.
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })
	// e.POST("/visit", visit)
	// e.Logger.Fatal(e.Start(":1323"))

}

func startServe() {
	s = v1.EchoHandler()
	s.Start(":1323")
}

func shutdownServer() {
	s.Shutdown(context.Background())
}
