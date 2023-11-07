package gorm

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(conn *sql.DB) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})
	return gormDB, err
}
