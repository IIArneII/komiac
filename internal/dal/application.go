package dal

import (
	"context"
	db_sql "database/sql"
	"errors"
	"komiac/internal/app/entities"
	app_errors "komiac/internal/app/errors"
	"komiac/internal/dal/models"
	"komiac/internal/dal/sql"
	"time"

	uuid "github.com/satori/go.uuid"
)

func (s *Storage) GetList(cxt context.Context, filter entities.ApplicationFilter) ([]*entities.Application, error) {
	return nil, errors.New("not implemented")
}

func (s *Storage) Create(cxt context.Context, application *entities.Application) (*entities.Application, error) {
	_, err := s.db.NamedExecContext(cxt, sql.ApplicationCreateSql, models.Application{
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
		CreatedAt:              time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return s.Get(cxt, application.UUID)
}

func (s *Storage) Get(cxt context.Context, uuid uuid.UUID) (*entities.Application, error) {
	applicationModel := models.Application{}

	err := s.db.NamedGetContext(cxt, &applicationModel, sql.ApplicationGetSql, sql.ApplicationGetParams{
		UUID: uuid,
	})
	if err == db_sql.ErrNoRows {
		s.log.WithField("uuid", uuid.String()).Info("Application not found")
		return nil, app_errors.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &entities.Application{
		UUID:                   applicationModel.UUID,
		MedicalOrganizationOID: applicationModel.MedicalOrganizationOID,
		DivisionOID:            applicationModel.DivisionOID,
		Year:                   applicationModel.Year,
		SMNN:                   applicationModel.SMNN,
		MNN:                    applicationModel.MNN,
		Form:                   applicationModel.Form,
		Dosage:                 applicationModel.Dosage,
		ConsumerUnit:           applicationModel.ConsumerUnit,
		ItemUnit:               applicationModel.ItemUnit,
		PrivilegeProgramCode:   applicationModel.PrivilegeProgramCode,
		PrivilegeProgram:       applicationModel.PrivilegeProgram,
	}, nil
}

func (s *Storage) Delete(cxt context.Context, uuid uuid.UUID) error {
	return errors.New("not implemented")
}

func (s *Storage) Update(cxt context.Context, application *entities.Application) (*entities.Application, error) {
	_, err := s.db.NamedExecContext(cxt, sql.ApplicationUpdateSql, models.Application{
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
	})
	if err != nil {
		return nil, err
	}

	return s.Get(cxt, application.UUID)
}
