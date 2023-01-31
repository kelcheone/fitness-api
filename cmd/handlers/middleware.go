package handlers

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

func LogRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		stop := time.Now()
		fmt.Printf("Request: %s %s %s %s\n", c.Request().Method, c.Request().URL, stop.Sub(start), c.Response().Status)
		return err
	}
}
