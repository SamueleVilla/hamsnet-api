package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PsqlConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func NewPsql(params PsqlConfig) (*sqlx.DB, error) {
	// connect to the database

	datasourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		params.Host, params.Port, params.User, params.Password, params.DBName)
	return sqlx.Connect("postgres", datasourceName)
}
