package sql

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(driver, connStr string) (repo *SQLRepository, err error) {
	sqlDB, err := sql.Open(driver, connStr)

	return &SQLRepository{
		db: sqlDB,
	}, err
}
