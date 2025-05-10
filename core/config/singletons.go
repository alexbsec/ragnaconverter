package config

import "github.com/alexbsec/ragconverter/api"

type Config struct {
	Env Environment
}

func GetRequester() api.DivineRequesterInterface {
	return api.NewDivineRequester()
}
