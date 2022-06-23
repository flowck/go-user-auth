package infra

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbDriver                  string `envconfig:"DB_DRIVER"`
	DbUrl                     string `envconfig:"DB_URL"`
	Port                      int16  `envconfig:"PORT"`
	Environment               string `envconfig:"AUTH_SERVICE_ENV"`
	JwtSigningKey             []byte `envconfig:"JWT_SIGNING_SECRET"`
	MaxRequestsPerMinutePerIp int    `envconfig:"MAX_REQUESTS_PER_MINUTE_PER_IP"`
}

var Configs *Config

func init() {
	Configs = &Config{}

	err := envconfig.Process("", Configs)

	if err != nil {
		panic(err)
	}

	fmt.Println("Environment", Configs.Environment)
	if Configs.Environment != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
}
