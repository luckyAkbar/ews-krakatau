package console

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "EWS Krakatau",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Panic(err.Error())

		os.Exit(1)
	}
}
