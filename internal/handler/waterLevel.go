package handler

import (
	"ews-krakatau/internal/constant"
	"ews-krakatau/internal/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func WaterLevel(c echo.Context) error {
	locationName := strings.ToLower(c.Param("locationName"))
	num, err := util.GetWaterLevelRecordsLength(locationName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, GenericError{
			OK:      false,
			Message: "Internal Server Error",
		})
	}

	if err := util.ValidateLocationName(locationName, constant.WATER_LEVEL_LOCATION); err != nil {
		return c.JSON(http.StatusBadRequest, GenericError{
			OK:      false,
			Message: fmt.Sprintf("Lokasi: %s tidak valid untuk water level sensor", locationName),
		})
	}

	randNum := util.RandNum(1, num)
	data, err := util.GetWaterLevelSensorData(randNum, locationName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, GenericError{
			OK:      false,
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, WaterLevelResult{
		OK:          true,
		LevRAD:      data.LevRAD,
		SensorTemp:  data.SensorTemp,
		Temperature: data.Temperature,
		PlaceName:   data.PlaceName,
	})
}
