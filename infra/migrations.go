package infra

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
	"os"
)

func RunMigrations(Db *sql.DB, config *Config) error {
	err := goose.SetDialect(config.DbDriver)

	workdir, _ := os.Getwd()

	err = goose.Up(Db, fmt.Sprintf("%s/sql", workdir))

	return err
}
