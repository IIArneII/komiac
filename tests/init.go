package tests

import (
	"database/sql"
	"komiac/internal/dal"
	"komiac/internal/def"

	"github.com/jmoiron/sqlx"
	"github.com/powerman/pqx"
	"github.com/powerman/sqlxx"
	"github.com/sirupsen/logrus"
)

var (
	log     = logrus.New()
	Storage = CreateStorage()
)

func connectDB(cfg pqx.Config) (*sqlxx.DB, error) {
	cfg.SSLMode = pqx.SSLDisable
	cfg.DefaultTransactionIsolation = sql.LevelSerializable

	db, err := sql.Open("pqx", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return sqlxx.NewDB(sqlx.NewDb(db, "postgres")), nil
}

func CreateStorage() *dal.Storage {
	cfg := pqx.Config{
		DBName: def.DBName,
		Host:   def.DBHost,
		Port:   def.DBPort,
		User:   def.DBUser,
		Pass:   def.DBPass,
	}

	db, err := connectDB(cfg)
	if err != nil {
		log.WithError(err).Fatal("DB is not connected")
		return nil
	}

	return dal.NewStorage(db)
}
