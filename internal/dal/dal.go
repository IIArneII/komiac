package dal

import (
	"github.com/powerman/sqlxx"
)

type Storage struct {
	db *sqlxx.DB
}

func NewStorage(db *sqlxx.DB) *Storage {
	return &Storage{
		db: db,
	}
}
