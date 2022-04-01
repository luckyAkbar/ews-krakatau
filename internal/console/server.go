package console

import (
	"ews-krakatau/internal/db"
	"ews-krakatau/internal/router"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var runServer = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long:  "Use this command to start EWS Krakatau HTTP server",
	Run:   InitServer,
}

func init() {
	RootCmd.AddCommand(runServer)
}

func InitServer(cmd *cobra.Command, args []string) {
	db.Connect()

	r := router.Router()
	s := http.Server{
		Addr:    ":3333",
		Handler: r,
	}

	log.Println("Server listening on port 3333")

	if err := s.ListenAndServe(); err != nil {
		log.Fatal("Failed to start the server")
	}
}
