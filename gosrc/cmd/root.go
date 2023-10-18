package cmd

import (
	"chatrooms/gosrc/cmd/server"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chatrooms",
	Short: "chatrooms app CLI",
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the application",
	Run:   server.Start,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
