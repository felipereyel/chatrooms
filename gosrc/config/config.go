package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type tconfigs struct {
	PublicDir          string
	ServerAddress      string
	DatabaseConnString string
	AutoMigrate        bool
	MigrationsDir      string
	JWTSecret          string
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

	envJWTSecret := os.Getenv("JWT_SECRET")
	if envJWTSecret != "" {
		Configs.JWTSecret = envJWTSecret
	} else {
		panic("JWT_SECRET is not set")
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
