package handler

import (
	"ews-krakatau/internal/constant"
	"ews-krakatau/internal/util"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Buoy(c echo.Context) error {
	locationName := strings.ToLower(c.Param("locationName"))
	num, err := util.GetBuoyRecordsLength(locationName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, GenericError{
			OK:      false,
			Message: "Internal Server Error",
		})
	}

	if err := util.ValidateLocationName(locationName, constant.BUOY_LOCATION); err != nil {
		return c.JSON(http.StatusBadRequest, GenericError{
			OK:      false,
			Message: fmt.Sprintf("Lokasi: %s tidak valid untuk buoy sensor", locationName),
		})
	}

	randNum := util.RandNum(1, num)
	data, err := util.GetBuoySensorData(randNum, locationName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, GenericError{
			OK:      false,
			Message: "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, BuoyResult{
		OK:         true,
		WaveHeight: data.WaveHeight,
		Latitude:   data.Latitude,
		Longitude:  data.Longitude,
	})
}
