package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type tconfigs struct {
	PublicDir          string
	ServerAddress      string
	RabbitMQConnString string
	DatabaseConnString string
	AutoMigrate        bool
	MigrationsDir      string
	JWTSecret          string
	BotUsername        string
	BotPassword        string
}

var Configs tconfigs

func init() {
	Configs.DatabaseConnString = os.Getenv("DATABASE_CONN_STRING")
	Configs.RabbitMQConnString = os.Getenv("RABBITMQ_CONN_STRING")
	Configs.BotUsername = os.Getenv("BOT_USERNAME")
	Configs.BotPassword = os.Getenv("BOT_PASSWORD")
	Configs.JWTSecret = os.Getenv("JWT_SECRET")

	// With default values
	publicDir := os.Getenv("PUBLIC_DIR")
	if publicDir != "" {
		Configs.PublicDir = publicDir
	} else {
		Configs.PublicDir = "./dist"
	}

	envPort := os.Getenv("PORT")
	if envPort != "" {
		Configs.ServerAddress = ":" + envPort
	} else {
		Configs.ServerAddress = ":3000"
	}

	envAutoMigrate := os.Getenv("AUTO_MIGRATE")
	if envAutoMigrate != "" {
		Configs.AutoMigrate = envAutoMigrate == "true"
	} else {
		Configs.AutoMigrate = true
	}

	envMigrationsDir := os.Getenv("MIGRATIONS_DIR")
	if envMigrationsDir != "" {
		Configs.MigrationsDir = envMigrationsDir
	} else {
		Configs.MigrationsDir = "migrations"
	}
}
