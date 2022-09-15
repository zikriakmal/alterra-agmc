package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HelloController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func StatusController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}
