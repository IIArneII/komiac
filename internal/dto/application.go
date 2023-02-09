package dto

import (
	api_models "komiac/internal/api/restapi/models"
	"komiac/internal/app/entities"
	dal_models "komiac/internal/dal/models"

	"github.com/go-openapi/strfmt"

	uuid "github.com/satori/go.uuid"
)

func AppToApiApplication(application *entities.Application) *api_models.Application {
	uuid := strfmt.UUID(application.UUID.String())
	year := int64(application.Year)
	return &api_models.Application{
		MNN:                    &application.MNN,
		SMNN:                   &application.SMNN,
		UUID:                   &uuid,
		ConsumerUnit:           &application.ConsumerUnit,
		DivisionOID:            &application.DivisionOID,
		Dosage:                 &application.Dosage,
		Form:                   &application.Form,
		ItemUnit:               &application.ItemUnit,
		MedicalOrganizationOID: &application.MedicalOrganizationOID,
		PrivilegeProgram:       &application.PrivilegeProgram,
		PrivilegeProgramCode:   &application.PrivilegeProgramCode,
		Year:                   &year,
	}
}

func ApiToAppApplication(application *api_models.Application) (*entities.Application, error) {
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

func DalToAppApplication(application dal_models.Application) *entities.Application {
	return &entities.Application{
		UUID:                   application.UUID,
		MedicalOrganizationOID: application.MedicalOrganizationOID,
		DivisionOID:            application.DivisionOID,
		Year:                   application.Year,
		SMNN:                   application.SMNN,
		MNN:                    application.MNN,
		Form:                   application.Form,
		Dosage:                 application.Dosage,
		ConsumerUnit:           application.ConsumerUnit,
		ItemUnit:               application.ItemUnit,
		PrivilegeProgramCode:   application.PrivilegeProgramCode,
		PrivilegeProgram:       application.PrivilegeProgram,
	}
}

func AppToApiApplications(applications []*entities.Application) api_models.Applications {
	applicationsApi := api_models.Applications{}
	for _, v := range applications {
		applicationsApi = append(applicationsApi, AppToApiApplication(v))
	}

	return applicationsApi
}

func ApiToAppApplications(applications api_models.Applications) ([]*entities.Application, error) {
	applicationsApp := make([]*entities.Application, 0, len(applications))

	for _, v := range applications {
		applicationApp, err := ApiToAppApplication(v)
		if err != nil {
			return nil, err
		}

		applicationsApp = append(applicationsApp, applicationApp)
	}

	return applicationsApp, nil
}

func DalToAppApplications(applications []dal_models.Application) []*entities.Application {
	applicationsApp := []*entities.Application{}
	for _, v := range applications {
		applicationsApp = append(applicationsApp, DalToAppApplication(v))
	}
	return applicationsApp
}
