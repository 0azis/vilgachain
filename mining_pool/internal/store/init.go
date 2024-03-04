package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitStore() (*sqlx.DB, error) {
	return sqlx.Connect("sqlite3", "/home/oazis/dns.db")
}
