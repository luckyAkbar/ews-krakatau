package models

import (
	"github.com/jinzhu/gorm"
)

type Buoy struct {
	gorm.Model

	WaveHeight int    `json:"wave_height"`
	Latitude   string `json:"latitude"`
	Longitude  string `json:"longitude"`
	PlaceName  string `json:"place_name"`
}
