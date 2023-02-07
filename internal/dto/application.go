package dto

import (
	"komiac/internal/api/restapi/models"
	"komiac/internal/app/entities"

	uuid "github.com/satori/go.uuid"
)

func ApiApplication(application entities.Application) models.Application {
	return models.Application{}
}

func AppApplication(application *models.Application) (*entities.Application, error) {
	uuid, err := uuid.FromString(application.UUID.String())
	if err != err {
		return nil, err
	}

	return &entities.Application{
		UUID:                   uuid,
		MedicalOrganizationOID: *application.MedicalOrganizationOID,
		DivisionOID:            *application.DivisionOID,
		Year:                   int(*application.Year),
		SMNN:                   *application.SMNN,
		MNN:                    *application.MNN,
		Form:                   *application.Form,
		Dosage:                 *application.Dosage,
		ConsumerUnit:           *application.ConsumerUnit,
		ItemUnit:               *application.ItemUnit,
		PrivilegeProgramCode:   *application.PrivilegeProgramCode,
		PrivilegeProgram:       *application.PrivilegeProgram,
	}, nil
}

func ApiApplications(applications []entities.Application) *models.Applications {
	return &models.Applications{}
}

func AppApplications(applications models.Applications) ([]*entities.Application, error) {
	applicationsApp := make([]*entities.Application, 0, len(applications))

	for _, v := range applications {
		applicationApp, err := AppApplication(v)
		if err != nil {
			return nil, err
		}

		applicationsApp = append(applicationsApp, applicationApp)
	}

	return applicationsApp, nil
}
