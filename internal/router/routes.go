package router

import (
	"ews-krakatau/internal/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type OK struct {
	Message string `json:"message"`
}

func Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/weather/:locationName", handler.Weather)
	e.GET("/seismic/:locationName", handler.Seismic)
	e.GET("/buoy/:locationName", handler.Buoy)
	e.GET("/water-level/:locationName", handler.WaterLevel)

	return e
}
