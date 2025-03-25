package api

import (
	"github.com/epuerta/callguard/go-backend/internal/middleware"
	"github.com/labstack/echo/v4"
)

// authMiddleware is an authentication middleware for Echo
func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return middleware.AuthEcho(c, next)
	}
}
