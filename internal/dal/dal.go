package dal

import (
	"github.com/powerman/sqlxx"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	db  *sqlxx.DB
	log *logrus.Logger
}

func NewStorage(db *sqlxx.DB) *Storage {
	return &Storage{
		db:  db,
		log: logrus.New(),
	}
}
