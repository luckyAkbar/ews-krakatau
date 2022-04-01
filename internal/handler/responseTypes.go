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
