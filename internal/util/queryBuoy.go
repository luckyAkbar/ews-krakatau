package util

import (
	"errors"
	"ews-krakatau/internal/db"
	"ews-krakatau/internal/models"
)

func GetBuoySensorData(index int, locationName string) (models.Buoy, error) {
	buoyData := models.Buoy{}

	err := db.DB.Model(&models.Buoy{}).
		Limit(1).
		Offset(index-1).
		Where("place_name = ?", locationName).
		Find(&buoyData).
		Error

	if err != nil {
		return buoyData, errors.New("Failed to get buoy data")
	}

	return buoyData, nil
}
