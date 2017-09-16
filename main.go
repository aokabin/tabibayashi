package main

import (
	"context"

	"github.com/aokabin/tabibayashi/kvs"
	"github.com/aokabin/tabibayashi/v1"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
)

var (
	s *echo.Echo
	r *redis.Client
)

func main() {
	defer shutdownServer()
	initDBClients()
	startServer()
}

func initDBClients() {
	r = kvs.KVSConnection()
}

func startServer() {
	s = v1.EchoHandler()
	s.Start(":1323")
}

func shutdownServer() {
	s.Shutdown(context.Background())
}
