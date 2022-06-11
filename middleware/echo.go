package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigureEcho() *echo.Echo {
	e := echo.New()
	xorm := ConfigureDatabase()
	//e.Validator = RegisterValidator()
	e.HideBanner = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())


	e.Use(middleware.RequestID())
	//e.Use(middleware.JWTWithConfig(JWT()))
	//e.Use(setSession())
	//e.Use(setLogger())
	e.Use(setXormSession(xorm))


	return e
}
