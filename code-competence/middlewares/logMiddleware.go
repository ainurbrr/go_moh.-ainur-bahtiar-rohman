package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LogMiddleware(e *echo.Echo){
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=[${time_rfc3339}], requestID=${id}, status=${status}, method=${method}, host=${host}, path=${path}, latency_human=${latency_human}, error=${error}\n",
	}))
}