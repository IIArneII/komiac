package main

import (
	"database/sql"
	"komiac/internal/api"
	"komiac/internal/app"
	"komiac/internal/dal"
	"komiac/internal/def"

	"github.com/jmoiron/sqlx"
	"github.com/powerman/pqx"
	"github.com/powerman/sqlxx"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New().WithField("module", "MAIN")
	cfg struct {
		gooseDir string
		db       pqx.Config
		api      api.Config
	}
)

func init() {
	cfg.api = api.Config{
		Host: def.Host,
		Port: def.Port,
	}
	cfg.db = pqx.Config{
		DBName: def.DBName,
		Host:   def.DBHost,
		Port:   def.DBPort,
		User:   def.DBUser,
		Pass:   def.DBPass,
	}
	cfg.gooseDir = def.GooseDir
}

func RunServer(errc chan<- error, app *app.App, cfg api.Config) {
	log.Info("Server started")
	defer log.Info("Server finished")

	server, err := api.NewServer(app, cfg)
	if err != nil {
		errc <- err
		return
	}
	errc <- server.Serve()
}

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

func migrateDB(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	if err := goose.Up(db, dir); err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := connectDB(cfg.db)
	if err != nil {
		log.WithError(err).Error("DB is not connected")
		return
	}

	err = migrateDB(db.DB.DB, cfg.gooseDir)
	if err != nil {
		log.WithError(err).Error("Migration error")
		return
	}

	repo := dal.NewStorage(db)

	app := app.NewApp(repo)

	errc := make(chan error)
	go RunServer(errc, app, cfg.api)

	err = <-errc
	if err != nil {
		log.Println(err)
	}
}
