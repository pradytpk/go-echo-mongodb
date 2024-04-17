package middlewares

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// ServerMessage acts has middleware
func ServerMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside middleware")
		return next(c)
	}
}
