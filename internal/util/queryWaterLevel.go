package util

import (
	"errors"
	"ews-krakatau/internal/db"
	"ews-krakatau/internal/models"
)

func GetWaterLevelSensorData(index int, locationName string) (models.WaterLevel, error) {
	waterLevel := models.WaterLevel{}

	err := db.DB.Model(&models.WaterLevel{}).
		Limit(1).
		Offset(index-1).
		Where("place_name = ?", locationName).
		Find(&waterLevel).
		Error

	if err != nil {
		return waterLevel, errors.New("Failed to get water level data")
	}

	return waterLevel, nil
}
