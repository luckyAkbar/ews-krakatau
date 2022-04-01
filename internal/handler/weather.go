package handler

import (
	"ews-krakatau/internal/constant"
	"ews-krakatau/internal/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Weather(c echo.Context) error {
	locationName := strings.ToLower(c.Param("locationName"))
	num, err := util.GetWeatherRecordsLength(locationName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, GenericError{
			OK:      false,
			Message: "Internal Server Error",
		})
	}

	if err := util.ValidateLocationName(locationName, constant.WEATHER_LOCATION); err != nil {
		return c.JSON(http.StatusBadRequest, GenericError{
			OK:      false,
			Message: fmt.Sprintf("Lokasi: %s tidak valid untuk weather sensor", locationName),
		})
	}

	randNum := util.RandNum(1, num)
	data, err := util.GetWeatherSensorData(randNum, locationName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, GenericError{
			OK:      false,
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, WeatherResult{
		OK:          true,
		WindSpeed:   data.WindSpeed,
		WindDir:     data.WindSpeed,
		Humidity:    data.Humidity,
		Temperature: data.Temperature,
		PlaceName:   data.PlaceName,
	})
}
