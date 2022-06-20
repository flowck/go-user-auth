package infra

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbDriver    string `envconfig:"DB_DRIVER"`
	DbUrl       string `envconfig:"DB_URL"`
	Port        int16  `envconfig:"PORT"`
	Environment string `envconfig:"AUTH_SERVICE_ENV"`
}

var Configs Config

func init() {
	Configs = Config{}

	err := envconfig.Process("", &Configs)

	if err != nil {
		panic(err)
		return
	}

	fmt.Println("Environment", Configs.Environment)
	if Configs.Environment != "production" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}
}
