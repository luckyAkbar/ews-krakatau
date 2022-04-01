package util

import (
	"errors"
	"ews-krakatau/internal/db"
	"ews-krakatau/internal/models"
)

func GetSeismicSensorData(index int, locationName string) (models.Seismic, error) {
	seismicData := models.Seismic{}

	err := db.DB.Model(&models.Seismic{}).
		Limit(1).
		Offset(index-1).
		Where("place_name = ?", locationName).
		Find(&seismicData).
		Error

	if err != nil {
		return seismicData, errors.New("Failed to get buoy data")
	}

	return seismicData, nil
}
