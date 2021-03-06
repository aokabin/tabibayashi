package v1

import (
	"github.com/labstack/echo"
)

func useV1Group(g *echo.Group) {
	g.POST("/visit", Visit)
	g.GET("/beacons", GetBeacons)
	g.POST("/music", CreateMusic)
	g.POST("/music_url", GetMusicURL)
}
