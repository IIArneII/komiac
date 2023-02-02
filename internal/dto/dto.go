package dto

import (
	"komiac/internal/api/restapi/models"
	"komiac/internal/app/entities"
)

func ApiMedicament(medicament *entities.Medicament) *models.Medicament {
	return &models.Medicament{
		SMNN:         medicament.SMNN,
		MNN:          medicament.MNN,
		ConsumerUnit: medicament.ConsumerUnit,
		Dosage:       medicament.Dosage,
		Form:         medicament.Form,
	}
}
