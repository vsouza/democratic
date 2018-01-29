package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/vsouza/democratic/config"
	"github.com/vsouza/democratic/handlers"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	enviroment = kingpin.Flag("environment", "Application environment").Default("dev").String()
)

func main() {
	kingpin.Parse()
	config.Init(*enviroment)
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/graphql", handlers.Query)
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
	// Start server
	c := config.GetConfig()
	e.Logger.Fatal(e.Start(c.GetString("server.port")))
}
