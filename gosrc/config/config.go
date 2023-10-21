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

	envDatabaseConnString := os.Getenv("DATABASE_CONN_STRING")
	if envDatabaseConnString != "" {
		Configs.DatabaseConnString = envDatabaseConnString
	} else {
		panic("DATABASE_CONN_STRING is not set")
	}

	envRabbitMQConnString := os.Getenv("RABBITMQ_CONN_STRING")
	if envRabbitMQConnString != "" {
		Configs.RabbitMQConnString = envRabbitMQConnString
	} else {
		panic("RABBITMQ_CONN_STRING is not set")
	}

	envJWTSecret := os.Getenv("JWT_SECRET")
	if envJWTSecret != "" {
		Configs.JWTSecret = envJWTSecret
	} else {
		panic("JWT_SECRET is not set")
	}

	envBotUsername := os.Getenv("BOT_USERNAME")
	if envBotUsername != "" {
		Configs.BotUsername = envBotUsername
	} else {
		panic("BOT_USERNAME is not set")
	}

	envBotPassword := os.Getenv("BOT_PASSWORD")
	if envBotPassword != "" {
		Configs.BotPassword = envBotPassword
	} else {
		panic("BOT_PASSWORD is not set")
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
