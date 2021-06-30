package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetMainMiddlewares(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Pre(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20.0)))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:  1 << 10, // 1 KB
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339} ${remote_ip} ${status} ${method} ${host}${path} ${latency_human}]\n",
	}))
}