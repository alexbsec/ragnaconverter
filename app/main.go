package main

import (
	"github.com/alexbsec/ragconverter/config"
	"github.com/alexbsec/ragconverter/logger"
)

func main() {
	log := logger.NewConsoleLogger()

	env, err := config.LoadEnv()
	if err != nil {
		panic("Could not load .env file!")
	}

}
