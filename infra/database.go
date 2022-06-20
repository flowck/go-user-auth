package infra

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDatabase(config *Config) error {
	var err error

	DB, err = sql.Open(config.DbDriver, config.DbUrl)

	return err
}
