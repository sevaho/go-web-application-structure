package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var Config ConfigType

type ConfigType struct {
	ENVIRONMENT               string `required:"True"`
	RELEASE                   string `default:"0.0.1"`
	DEVELOPMENT               bool   `default:"true"`
	HTTP_SERVER_PORT          int    `default:"3000"`
	DB_DSN                    string `required:"True"`
	DB_CONNECTION_RETRIES     int    `default:"100"`
	DB_CONNECTION_RETRY_DELAY time.Duration
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func init() {
	Config.DB_CONNECTION_RETRY_DELAY = 2 * time.Second

	// load .env file to env vars if any
	_ = godotenv.Load()

	// parse env vars to struct
	err := envconfig.Process("", &Config)
	if err != nil {
		log.Fatalf("Failed to decode env vars to struct config/config.go: %s", err)
	}

	// set development off if running remote
	if stringInSlice(Config.ENVIRONMENT, [](string){"QA", "STAGING", "PRODUCTION"}) {
		Config.DEVELOPMENT = false
	}
}
