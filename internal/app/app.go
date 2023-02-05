package app

import "komiac/internal/dal"

type App struct {
	ApplicationSvc ApplicationService
}

func NewApp(storage *dal.Storage) *App {
	return &App{
		ApplicationSvc: NewApplicationService(storage),
	}
}
