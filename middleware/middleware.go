package middleware

import (
	"fmt"

	"github.com/labstack/echo"
)

func V1GroupFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("Before")
		next(c)
		fmt.Println("After")
		return nil
	}
}
