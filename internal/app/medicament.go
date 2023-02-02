package app

import (
	"komiac/internal/app/entities"
)

type MedicamentService interface {
	Get(SMNN string) (*entities.Medicament, error)
	Add(entities.Medicament) (*entities.Medicament, error)
}

type MedicamentRepository interface {
	Get(SMNN string) (*entities.Medicament, error)
	Add(entities.Medicament) (*entities.Medicament, error)
}

type medicamentSvc struct {
	repo MedicamentRepository
}

func NewMedicamentService(repo MedicamentRepository) MedicamentService {
	return &medicamentSvc{
		repo: repo,
	}
}

func (svc *medicamentSvc) Get(SMNN string) (*entities.Medicament, error) {
	return nil, nil
}

func (svc *medicamentSvc) Add(entities.Medicament) (*entities.Medicament, error) {
	return nil, nil
}
