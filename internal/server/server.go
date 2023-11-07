package app

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo/v4"
	log2 "github.com/labstack/gommon/log"
	"github.com/rs/zerolog/log"

	v1 "crud-app/internal/api/controller/v1"
	"crud-app/internal/config"
	"crud-app/internal/domain"
	"crud-app/internal/repository/person_repository"
	"crud-app/pkg/database/gorm"
	"crud-app/pkg/logger"
)

func NewPgSQLConnection(conn *config.PostgresConfig) *sql.DB {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conn.User, conn.Password,
		conn.Host, conn.Port, conn.Database)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func Run(cfg *config.Config, dbCfg *config.PostgresConfig) {
	l := logger.New(cfg.Level)
	l.Debug(dbCfg)
	db := NewPgSQLConnection(dbCfg)
	defer db.Close()

	orm, err := gorm.New(db)
	if err != nil {
		panic(err)
	}
	if err = orm.AutoMigrate(&domain.Person{}); err != nil {
		log.Print(err)
	}
	repo := person_repository.New(orm)
	handler := echo.New()
	handler.Logger.SetLevel(log2.DEBUG)
	v1.NewRouter(handler, repo)
	l.Fatal(handler.Start(fmt.Sprintf(":%s", cfg.Port)))
}
