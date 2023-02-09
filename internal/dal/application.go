package dal

import (
	"context"
	db_sql "database/sql"
	"komiac/internal/app/entities"
	app_errors "komiac/internal/app/errors"
	"komiac/internal/dal/models"
	"komiac/internal/dal/sql"
	"time"

	uuid "github.com/satori/go.uuid"
)

func (s *Storage) GetList(cxt context.Context, filter entities.ApplicationFilter) ([]*entities.Application, error) {
	applicationModel := []models.Application{}

	err := s.db.NamedSelectContext(cxt, &applicationModel, sql.AplicationGetListSql, sql.ApplicationGetListParams{
		DivisionOid: *filter.DivisionOID,
		Year:        *filter.Year,
		Mnn:         *filter.MNN,
	})

	if err == db_sql.ErrNoRows {
		return nil, app_errors.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return appLists(applicationModel), nil
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
	t := time.Now()
	res, err := s.db.NamedExec(sql.AplicationDeleteSql, sql.ApplicationDeleteParams{
		UUID:      uuid,
		DeletedAt: t,
	})
	if err != nil {
		return err
	}
	if count, _ := res.RowsAffected(); count == 0 {
		return app_errors.ErrNotFound
	}

	return nil
}

func (s *Storage) Update(cxt context.Context, application *entities.Application) (*entities.Application, error) {
	modifiedAt := time.Now()
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
		ModifiedAt:             &modifiedAt,
	})
	if err != nil {
		return nil, err
	}

	return s.Get(cxt, application.UUID)
}

func appList(m models.Application) *entities.Application {
	return &entities.Application{
		UUID:                   m.UUID,
		MedicalOrganizationOID: m.MedicalOrganizationOID,
		DivisionOID:            m.DivisionOID,
		Year:                   m.Year,
		SMNN:                   m.SMNN,
		MNN:                    m.MNN,
		Form:                   m.Form,
		Dosage:                 m.Dosage,
		ConsumerUnit:           m.ConsumerUnit,
		ItemUnit:               m.ItemUnit,
		PrivilegeProgramCode:   m.PrivilegeProgramCode,
		PrivilegeProgram:       m.PrivilegeProgram,
	}
}

func appLists(ms []models.Application) []*entities.Application {
	ams := []*entities.Application{}
	for _, m := range ms {
		ams = append(ams, appList(m))
	}
	return ams
}
