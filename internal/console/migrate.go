package console

import (
	"ews-krakatau/internal/db"
	"ews-krakatau/internal/models"
	"log"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate all database table",
	Long:  "Use this command to initialize your database table for the first time",
	Run:   migrate,
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}

func migrate(cmd *cobra.Command, args []string) {
	db.Connect()

	db.DB.AutoMigrate(&models.Buoy{})
	db.DB.AutoMigrate(&models.Weather{})
	db.DB.AutoMigrate(&models.Seismic{})
	db.DB.AutoMigrate(&models.WaterLevel{})

	log.Println("Migration finished")
}
