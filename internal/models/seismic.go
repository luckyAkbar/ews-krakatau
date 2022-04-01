package models

import (
	"gorm.io/gorm"
)

type Seismic struct {
	gorm.Model

	Magnitude float32 `json:"magnitude"`
	PlaceName string  `json:"place_name"`
}
