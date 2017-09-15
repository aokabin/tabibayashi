package v1

import (
	"github.com/labstack/echo"
)

func useSampleGroup(g *echo.Group) {
	g.POST("/visit", visit)
}
