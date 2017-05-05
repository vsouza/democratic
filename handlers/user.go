package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func getUserByID(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
