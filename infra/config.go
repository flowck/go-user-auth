package infra

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbDriver string `envconfig:"DB_DRIVER"`
	DbUrl    string `envconfig:"DB_URL"`
	Port     int16  `envconfig:"PORT"`
}

var Configs Config

func init() {
	Configs = Config{}

	if err := godotenv.Load(); err != nil {
		panic(err)
		return
	}

	err := envconfig.Process("", &Configs)

	if err != nil {
		panic(err)
	}
}
