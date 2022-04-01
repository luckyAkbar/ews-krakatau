package models

import (
	"gorm.io/gorm"
)

type WaterLevel struct {
	gorm.Model

	LevRAD      int    `json:"lev_rad"`
	SensorTemp  int    `json:"sensor_temp"`
	Temperature int    `json:"temperature"`
	PlaceName   string `json:"place_name"`
}
