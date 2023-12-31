package cmd

import (
	"chatrooms/gosrc/cmd/bot"
	"chatrooms/gosrc/cmd/migrate"
	"chatrooms/gosrc/cmd/server"
	"chatrooms/gosrc/cmd/tap"

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

var migrateUpCmd = &cobra.Command{
	Use:   "migrate:up",
	Short: "Migrates Up the database",
	Run:   migrate.Up,
}

var migrateDownCmd = &cobra.Command{
	Use:   "migrate:down",
	Short: "Migrates Down the database",
	Run:   migrate.Down,
	Args:  cobra.ExactArgs(1),
}

var botCmd = &cobra.Command{
	Use:   "bot",
	Short: "Run the bot",
	Run:   bot.Run,
}

var tapCmd = &cobra.Command{
	Use:   "tap",
	Short: "Tap a room or all rooms",
	Run:   tap.Tap,
	Args:  cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.AddCommand(tapCmd)
	rootCmd.AddCommand(botCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(migrateUpCmd)
	rootCmd.AddCommand(migrateDownCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err.Error())
	}
}
