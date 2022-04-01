package util

import (
	"errors"
	"ews-krakatau/internal/db"
	"ews-krakatau/internal/models"
)

func GetWeatherSensorData(index int, locationName string) (models.Weather, error) {
	weatherData := models.Weather{}

	err := db.DB.Model(&models.Weather{}).
		Limit(1).
		Offset(index-1).
		Where("place_name = ?", locationName).
		Find(&weatherData).
		Error

	if err != nil {
		return weatherData, errors.New("Failed to get buoy data")
	}

	return weatherData, nil
}
