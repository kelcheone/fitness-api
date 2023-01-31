package main

import (
	"fitness-api/cmd/handlers"
	"fitness-api/cmd/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.Home)
	storage.InitDB()

	// e.Use(handlers.LogRequest)
	e.Use(middleware.Logger(), middleware.Recover())

	e.POST("/users", handlers.CreateUser)
	e.POST("/measurements", handlers.CreateMeasurement)
	e.PUT("/user/:id", handlers.HandleUpdateUser)
	e.PUT("/measurement/:id", handlers.HandleUpdateMeasurement)

	e.Logger.Fatal(e.Start(":8080"))
}
