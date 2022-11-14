package internal

import (
	"log"
	"os"
)

const (
	baseUrl        = "OPENAI_API_BASE_URL"
	defaultVersion = "OPENAI_API_VERSION"
	apiKeyName     = "OPENAI_API_KEY"
)

var Config config

type config struct {
	Gpt3BaseUrl    string
	Gpt3ApiVersion string
	Gpt3ApiKey     string
}

func init() {
	Config.Gpt3BaseUrl = os.Getenv(baseUrl)
	if Config.Gpt3BaseUrl == "" {
		Config.Gpt3BaseUrl = "https://api.openai.com"
	}

	Config.Gpt3ApiVersion = os.Getenv(defaultVersion)
	if Config.Gpt3ApiVersion == "" {
		Config.Gpt3ApiVersion = "v1"
	}

	Config.Gpt3ApiKey = os.Getenv(apiKeyName)
	if Config.Gpt3ApiKey == "" {
		log.Panicf("Api key required. Please ensure env variable %s is set.", apiKeyName)
	}
}
