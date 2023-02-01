package app

import "komiac/internal/dal"

type App interface {
}

type app struct {
	storage dal.Storage
}

func NewApp(storage dal.Storage) App {
	return app{
		storage: storage,
	}
}
