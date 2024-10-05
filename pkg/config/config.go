package config

import (
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	JWTsecret   string
	ExpTime     int
	Port        string
	Addr        string
	Mongo       string
}

func New() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	time := os.Getenv("EXP")
	t, err := strconv.Atoi(time)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	return &Config{
		Environment: os.Getenv("ENV"),
		JWTsecret:   os.Getenv("JWT"),
		ExpTime:     t,
		Port:        os.Getenv("PORT"),
		Addr:        os.Getenv("ADDR"),
		Mongo:       os.Getenv("MONGO"),
	}
}
