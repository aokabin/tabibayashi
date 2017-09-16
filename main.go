package main

import (
	"context"

	"github.com/motemen/ghq/utils"

	"github.com/aokabin/tabibayashi/config"
	"github.com/aokabin/tabibayashi/kvs"
	"github.com/aokabin/tabibayashi/v1"
	"github.com/labstack/echo"
)

var (
	s *echo.Echo
)

func main() {
	defer shutdownServer()
	utils.PanicIf(kvs.Connection())
	startServer()
}

func startServer() {
	s = v1.EchoHandler()
	s.Start(":" + config.ServerPort())
}

func shutdownServer() {
	s.Shutdown(context.Background())
}
