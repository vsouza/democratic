package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", handlers.hello)
	e.GET("/:id", handlers.getUserByID)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
