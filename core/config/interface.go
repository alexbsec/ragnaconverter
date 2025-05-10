package config

import "github.com/alexbsec/ragconverter/api"

type ConfigInterface interface {
	GetRequester() api.DivineRequesterInterface
}
