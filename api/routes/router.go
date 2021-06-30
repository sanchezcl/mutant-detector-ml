package routes

import (
	"github.com/labstack/echo/v4"
	"mutantDetector/api/middlewares"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	//set api version groups
	apiV1 := e.Group("/v1")

	//set middlewares
	middlewares.SetMainMiddlewares(e)

	//setRoutes
	apiV1Routes(apiV1)

	return e
}
