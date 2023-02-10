package dal

import (
	"context"
	db_sql "database/sql"
	"komiac/internal/app/entities"
	app_errors "komiac/internal/app/errors"
	"komiac/internal/dal/models"
	"komiac/internal/dal/sql"
	"komiac/internal/dto"
	"time"

	_uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func (s *Storage) GetList(cxt context.Context, filter entities.ApplicationFilter) ([]*entities.Application, error) {
	applicationModels := []models.Application{}
	var filterStrs string
	namedVars := map[string]interface{}{}

	if filter.DivisionOID != nil {
		filterStrs += "AND division_oid=:DivisionOID "
		namedVars["DivisionOID"] = filter.DivisionOID
	}
	if filter.Year != nil {
		filterStrs += "AND year=:Year "
		namedVars["Year"] = filter.Year
	}
	if filter.MNN != nil {
		filterStrs += "AND mnn=:MNN "
		namedVars["MNN"] = filter.MNN
	}

	err := s.db.NamedSelectContext(cxt, &applicationModels, sql.AplicationGetListSql+filterStrs, namedVars)

	if err == db_sql.ErrNoRows {
		s.log.WithFields(logrus.Fields{
			"DivisionOID": filter.DivisionOID,
			"Year":        filter.Year,
			"MNN":         filter.MNN,
		}).Info("Applications not found")
		return []*entities.Application{}, nil
	} else if err != nil {
		return nil, err
	}

	return dto.DalToAppApplications(applicationModels), nil
}

func (s *Storage) Create(cxt context.Context, application *entities.Application) (*entities.Application, error) {
	if _uuid.Equal(application.UUID, _uuid.NullUUID{}.UUID) {
		return nil, app_errors.BadUUID
	}

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

func (s *Storage) Get(cxt context.Context, uuid _uuid.UUID) (*entities.Application, error) {
	if _uuid.Equal(uuid, _uuid.NullUUID{}.UUID) {
		return nil, app_errors.BadUUID
	}

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

func (s *Storage) Delete(cxt context.Context, uuid _uuid.UUID) error {
	result, err := s.db.NamedExec(sql.AplicationDeleteSql, sql.ApplicationDeleteParams{
		UUID:      uuid,
		DeletedAt: time.Now(),
	})
	if err != nil {
		return err
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return app_errors.ErrNotFound
	}

	return nil
}

func (s *Storage) Update(cxt context.Context, application *entities.Application) (*entities.Application, error) {
	if _uuid.Equal(application.UUID, _uuid.NullUUID{}.UUID) {
		return nil, app_errors.BadUUID
	}

	modifiedAt := time.Now()
	result, err := s.db.NamedExecContext(cxt, sql.ApplicationUpdateSql, models.Application{
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

	if count, _ := result.RowsAffected(); count == 0 {
		return nil, app_errors.ErrNotFound
	}

	return s.Get(cxt, application.UUID)
}
