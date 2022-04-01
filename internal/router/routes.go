package router

import (
	"ews-krakatau/internal/handler"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type OK struct {
	Message string `json:"message"`
}

func Router() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/weather/sertung", func(c echo.Context) error {
		return c.JSON(http.StatusOK, OK{
			Message: "oke",
		})
	})
	e.GET("/weather/rakata", func(c echo.Context) error {
		return c.JSON(http.StatusOK, OK{
			Message: "oke",
		})
	})

	e.GET("/seismic/sertung", func(c echo.Context) error {
		return c.JSON(http.StatusOK, OK{
			Message: "oke",
		})
	})
	e.GET("/seismic/anak-krakatau", func(c echo.Context) error {
		return c.JSON(http.StatusOK, OK{
			Message: "oke",
		})
	})
	e.GET("/seismic/panjang", func(c echo.Context) error {
		return c.JSON(http.StatusOK, OK{
			Message: "oke",
		})
	})
	e.GET("/seismic/rakata", func(c echo.Context) error {
		return c.JSON(http.StatusOK, OK{
			Message: "oke",
		})
	})

	e.GET("/buoy/:locationName", handler.Buoy)

	e.GET("/water-level/pematang", func(c echo.Context) error {
		return c.JSON(http.StatusOK, OK{
			Message: "oke",
		})
	})
	e.GET("/water-level/hanura", func(c echo.Context) error {
		return c.JSON(http.StatusOK, OK{
			Message: "oke",
		})
	})

	return e
}
