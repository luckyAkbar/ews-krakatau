package models

import (
	"github.com/jinzhu/gorm"
)

type Weather struct {
	gorm.Model

	WindSpeed     int    `json:"wind_speed"`
	WindDirection int    `json:"wind_direction"`
	Humidity      int    `json:"humidity"`
	Temperature   int    `json:"temperature"`
	PlaceName     string `json:"place_name"`
}
