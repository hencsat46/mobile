package config

import (
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Environment string `yaml:"env"`
	JWTsecret   string `yaml:"secretKey"`
	ExpTime     int    `yaml:"expTime"`
	Port        string `yaml:"port"`
	Addr        string `yaml:"host"`
	Mongo       string `yaml:"database"`
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

func NewYaml(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)

	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	if err := decoder.Decode(&config); err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	return config, nil
}
