package v1

import (
	swaggerFiles "github.com/swaggo/files"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"crud-app/internal/api/controller/handler/person"
)

func NewRouter(handler *echo.Echo, repo storage) {
	handler.Use(middleware.Logger())
	handler.Use(middleware.Recover())

	handler.GET("/swagger/*any", echo.WrapHandler(swaggerFiles.Handler))

	handler.GET("/healthz", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "ok",
		})
	})

	handler.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	// Routers
	g := handler.Group("/api/v1")
	{
		person.New(g, repo)
	}
}
