package config

import (
	"os"

	"github.com/alexbsec/ragconverter/logger"
	"github.com/joho/godotenv"
)

var log *logger.ConsoleLogger = logger.NewConsoleLogger()

type Environment struct {
	ApiKey string
	Host   string
}

func LoadEnv() (*Environment, error) {
	err := godotenv.Load()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	env := &Environment{
		ApiKey: os.Getenv("DIVINE_PRIDE_API_KEY"),
	}

	return env, nil
}
