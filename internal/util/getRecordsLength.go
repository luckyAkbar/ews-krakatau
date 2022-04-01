package util

import (
	"errors"
	"ews-krakatau/internal/db"
	"ews-krakatau/internal/models"
)

func GetBuoyRecordsLength(locationName string) (int, error) {
	var num int

	err := db.DB.Model(&models.Buoy{}).Select("COUNT(*)").Where("place_name = ?", locationName).Scan(&num).Error

	if err != nil {
		return 0, errors.New("Failed to count buoy records")
	}

	return num, nil
}
