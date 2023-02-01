package dal

import (
	"github.com/powerman/sqlxx"
)

type storage struct {
	db *sqlxx.DB
}

type Storage interface {
}

func NewRepo(db *sqlxx.DB) Storage {
	return storage{
		db: db,
	}
}
