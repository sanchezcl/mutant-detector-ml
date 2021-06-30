package routes

import (
	"github.com/labstack/echo/v4"
	"mutantDetector/api/controllers"
)

func apiV1Routes(group *echo.Group) {
	group.POST("/mutant", controllers.MutantDetector)
	group.GET("/stats", controllers.MutantStats)
}