package handler

type GenericError struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

type BuoyResult struct {
	OK         bool   `json:"ok"`
	WaveHeight int    `json:"wave_height"`
	Latitude   string `json:"lat"`
	Longitude  string `json:"long"`
}

type WeatherResult struct {
	OK          bool   `json:"ok"`
	WindSpeed   int    `json:"wind_speed"`
	WindDir     int    `json:"wind_direction"`
	Humidity    int    `json:"humidity"`
	Temperature int    `json:"temp"`
	PlaceName   string `json:"place_name"`
}

type SeismicResult struct {
	OK        bool    `json:"ok"`
	Magnitude float32 `json:"magnitude"`
	PlaceName string  `json:"place_name"`
}

type WaterLevelResult struct {
	OK          bool   `json:"ok"`
	LevRAD      int    `json:"lev_rad"`
	SensorTemp  int    `json:"sensor_temp"`
	Temperature int    `json:"temperature"`
	PlaceName   string `json:"place_name"`
}
