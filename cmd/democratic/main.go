package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/vsouza/Democratic/handlers"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/graphql", handlers.Query)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
