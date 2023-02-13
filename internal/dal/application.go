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
	namedVars := logrus.Fields{}

	if filter.DivisionOID != nil {
		filterStrs += "AND division_oid=:DivisionOID "
		namedVars["DivisionOID"] = filter.DivisionOID
	}
	if filter.Year != nil {
		filterStrs += "AND year=:Year "
		namedVars["Year"] = filter.Year
	}
	if filter.MNN != nil {
		filterStrs += "AND LOWER(mnn)=LOWER(:MNN) "
		namedVars["MNN"] = filter.MNN
	}

	err := s.db.NamedSelectContext(cxt, &applicationModels, sql.AplicationGetListSql+filterStrs, namedVars)

	if err == db_sql.ErrNoRows {
		s.log.WithFields(namedVars).Info("GetList: applications not found")
		return []*entities.Application{}, nil
	} else if err != nil {
		s.log.WithError(err).WithFields(namedVars).Error("GetList: erorr")
		return nil, err
	}

	return dto.DalToAppApplications(applicationModels), nil
}

func (s *Storage) Create(cxt context.Context, application *entities.Application) (*entities.Application, error) {
	if application == nil {
		s.log.WithError(app_errors.ErrEntityIsNil).Error("Create: erorr")
		return nil, app_errors.ErrEntityIsNil
	}

	if _uuid.Equal(application.UUID, _uuid.NullUUID{}.UUID) {
		s.log.WithError(app_errors.ErrBadUUID).Error("Create: erorr")
		return nil, app_errors.ErrBadUUID
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
		s.log.WithError(err).WithField("uuid", application.UUID.String()).Error("Create: erorr")
		return nil, err
	}

	return s.Get(cxt, application.UUID, false)
}

func (s *Storage) Get(cxt context.Context, uuid _uuid.UUID, withDeleted bool) (*entities.Application, error) {
	sqlGet := sql.ApplicationGetSql
	if withDeleted {
		sqlGet = sql.ApplicationGetWithDeletedSql
	}

	applicationModel := models.Application{}

	err := s.db.NamedGetContext(cxt, &applicationModel, sqlGet, sql.ApplicationGetParams{
		UUID: uuid,
	})
	if err == db_sql.ErrNoRows {
		s.log.WithField("uuid", uuid.String()).Info("Get: application not found")
		return nil, app_errors.ErrNotFound
	} else if err != nil {
		s.log.WithError(err).Error("Get: erorr")
		return nil, err
	}

	return dto.DalToAppApplication(applicationModel), nil
}

func (s *Storage) Delete(cxt context.Context, uuid _uuid.UUID) error {
	result, err := s.db.NamedExec(sql.AplicationDeleteSql, sql.ApplicationDeleteParams{
		UUID:      uuid,
		DeletedAt: time.Now(),
	})
	if err != nil {
		s.log.WithError(err).Error("Get: erorr")
		return err
	}

	if count, _ := result.RowsAffected(); count == 0 {
		s.log.WithField("uuid", uuid.String()).Info("Delete: application not found")
		return app_errors.ErrNotFound
	}

	return nil
}

func (s *Storage) Update(cxt context.Context, application *entities.Application) (*entities.Application, error) {
	if application == nil {
		s.log.WithError(app_errors.ErrEntityIsNil).Error("Update: erorr")
		return nil, app_errors.ErrEntityIsNil
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
		DeletedAt:              nil,
	})
	if err != nil {
		s.log.WithError(err).Error("Update: erorr")
		return nil, err
	}

	if count, _ := result.RowsAffected(); count == 0 {
		s.log.WithField("uuid", application.UUID.String()).Info("Update: application not found")
		return nil, app_errors.ErrNotFound
	}

	return s.Get(cxt, application.UUID, false)
}
