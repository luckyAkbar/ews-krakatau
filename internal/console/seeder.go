package console

import (
	"encoding/json"
	"ews-krakatau/internal/constant"
	"ews-krakatau/internal/db"
	"ews-krakatau/internal/models"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

type Location struct {
	Latitude  string `json:"lat"`
	Longitude string `json:"long"`
}

type WeatherData struct {
	WindSpeed []int    `json:"windSpeed"`
	WindDir   []int    `json:"windDir"`
	Humidity  []int    `json:"humidity"`
	Temp      []int    `json:"Temp"`
	Location  Location `json:"location"`
}

type SeismicData struct {
	Magnitude []float32 `json:"magnitude"`
	Location  Location  `json:"location"`
}

type BuoyData struct {
	WaveHeight []int    `json:"waveHeight"`
	Location   Location `json:"location"`
}

type WaterLevelData struct {
	LevRAD     []int    `json:"levRAD"`
	SensorTemp []int    `json:"sensorTemp"`
	Temp       []int    `json:"temp"`
	Location   Location `json:"location"`
}

type WeatherSensor struct {
	Rakata  WeatherData `json:"rakata"`
	Sertung WeatherData `json:"sertung"`
}

type SeismicSensor struct {
	Sertung      SeismicData `json:"sertung"`
	AnakKrakatau SeismicData `json:"gunungAnakKrakatau"`
	Panjang      SeismicData `json:"panjang"`
	Rakata       SeismicData `json:"rakata"`
}

type BuoySensor struct {
	Sertung BuoyData `json:"sertung"`
	Rakata  BuoyData `json:"rakata"`
}

type WaterLevelSensor struct {
	Pematang WaterLevelData `json:"pematang"`
	Hanura   WaterLevelData `json:"hanura"`
}

type Payload struct {
	Weather    WeatherSensor    `json:"weather"`
	Seismic    SeismicSensor    `json:"seismic"`
	Buoy       BuoySensor       `json:"buoy"`
	WaterLevel WaterLevelSensor `json:"waterLevel"`
}

var seederCmd = &cobra.Command{
	Use:   "seeder",
	Short: "Run the seeder program",
	Long:  "Use this command to seed your database",
	Run:   seeder,
}

func init() {
	RootCmd.AddCommand(seederCmd)
}

func seeder(cmd *cobra.Command, args []string) {
	db.Connect()

	payload := unmarshallPayload()

	fmt.Println("Start seeding database")

	seedBuoy(payload.Buoy)
	seedWeather(payload.Weather)
	seedSeismic(payload.Seismic)
	seedWaterLevel(payload.WaterLevel)

	fmt.Println("Finished.")
}

func seedWaterLevel(waterLevel WaterLevelSensor) {
	pematang := createWaterLevelData(waterLevel.Pematang, constant.PEMATANG)
	hanura := createWaterLevelData(waterLevel.Hanura, constant.HANURA)

	db.DB.Create(&pematang)
	db.DB.Create(&hanura)
}

func createWaterLevelData(data WaterLevelData, placeName string) []models.WaterLevel {
	records := []models.WaterLevel{}

	for i, record := range data.LevRAD {
		waterLevel := models.WaterLevel{
			LevRAD:      record,
			SensorTemp:  data.SensorTemp[i],
			Temperature: data.Temp[i],
			PlaceName:   placeName,
		}

		records = append(records, waterLevel)
	}

	return records
}

func seedSeismic(seismic SeismicSensor) {
	sertung := createSeismicData(seismic.Sertung, constant.SERTUNG)
	anakKrakatau := createSeismicData(seismic.AnakKrakatau, constant.ANAK_KRAKATAU)
	panjang := createSeismicData(seismic.Panjang, constant.PANJANG)
	rakata := createSeismicData(seismic.Rakata, constant.RAKATA)

	db.DB.Create(&rakata)
	db.DB.Create(&sertung)
	db.DB.Create(&anakKrakatau)
	db.DB.Create(&panjang)
}

func createSeismicData(data SeismicData, placeName string) []models.Seismic {
	records := []models.Seismic{}

	for _, record := range data.Magnitude {
		seismic := models.Seismic{
			Magnitude: record,
			PlaceName: placeName,
		}

		records = append(records, seismic)
	}

	return records
}

func seedWeather(weather WeatherSensor) {
	sertung := createWeatherData(weather.Sertung, constant.SERTUNG)
	rakata := createWeatherData(weather.Rakata, constant.RAKATA)

	db.DB.Create(&sertung)
	db.DB.Create(&rakata)
}

func createWeatherData(data WeatherData, placeName string) []models.Weather {
	records := []models.Weather{}

	for i, record := range data.WindSpeed {
		weather := models.Weather{
			WindSpeed:     record,
			WindDirection: data.WindDir[i],
			Humidity:      data.Humidity[i],
			Temperature:   data.Temp[i],
			PlaceName:     placeName,
		}

		records = append(records, weather)
	}

	return records
}

func seedBuoy(buoy BuoySensor) {
	sertung := createBuoyData(buoy.Sertung, constant.SERTUNG)
	rakata := createBuoyData(buoy.Rakata, constant.RAKATA)

	db.DB.Create(&sertung)
	db.DB.Create(&rakata)
}

func createBuoyData(data BuoyData, placeName string) []models.Buoy {
	records := []models.Buoy{}

	for _, record := range data.WaveHeight {
		buoy := models.Buoy{
			WaveHeight: record,
			Latitude:   data.Location.Latitude,
			Longitude:  data.Location.Longitude,
			PlaceName:  placeName,
		}

		records = append(records, buoy)
	}

	return records
}

func unmarshallPayload() Payload {
	JSONFile, err := os.Open("./data/data.json")

	if err != nil {
		log.Panicf("Failed to open JSON file: %s", err.Error())
	}

	defer JSONFile.Close()

	bytes, _ := ioutil.ReadAll(JSONFile)

	var payload Payload

	json.Unmarshal(bytes, &payload)

	return payload
}
