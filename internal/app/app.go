package app

import "komiac/internal/dal"

type App struct {
	MedicamentSvc MedicamentService
}

func NewApp(storage *dal.Storage) *App {
	return &App{
		MedicamentSvc: NewMedicamentService(storage),
	}
}
