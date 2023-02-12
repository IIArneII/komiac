package app

import (
	"context"
	"komiac/internal/app/entities"
	app_errors "komiac/internal/app/errors"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type ApplicationService interface {
	GetList(cxt context.Context, filter entities.ApplicationFilter) ([]*entities.Application, error)
	AddList(cxt context.Context, application []*entities.Application) error
	Delete(cxt context.Context, uuid uuid.UUID) error
}

type ApplicationRepository interface {
	GetList(cxt context.Context, filter entities.ApplicationFilter) ([]*entities.Application, error)
	Get(cxt context.Context, uuid uuid.UUID, withDeleted bool) (*entities.Application, error)
	Create(cxt context.Context, application *entities.Application) (*entities.Application, error)
	Delete(cxt context.Context, uuid uuid.UUID) error
	Update(cxt context.Context, application *entities.Application) (*entities.Application, error)
}

type ApplicationSvc struct {
	repo ApplicationRepository
	log  *logrus.Entry
}

func NewApplicationService(repo ApplicationRepository) ApplicationService {
	return &ApplicationSvc{
		repo: repo,
		log:  logrus.New().WithField("module", "APP"),
	}
}

func (svc *ApplicationSvc) GetList(cxt context.Context, filter entities.ApplicationFilter) ([]*entities.Application, error) {
	fields := logrus.Fields{}
	if filter.DivisionOID != nil {
		fields["DivisionOID"] = *filter.DivisionOID
	}
	if filter.MNN != nil {
		fields["MNN"] = *filter.MNN
	}
	if filter.Year != nil {
		fields["Year"] = *filter.Year
	}
	svc.log.WithFields(fields).Info("GetList")

	applications, err := svc.repo.GetList(cxt, filter)
	if err != nil {
		svc.log.WithError(err).WithFields(fields).Error("GetList: error")
		return nil, err
	}

	svc.log.WithField("count", len(applications)).Info("GetList: received applications list")

	return applications, nil
}

func (svc *ApplicationSvc) AddList(cxt context.Context, applications []*entities.Application) error {
	svc.log.WithField("count", len(applications)).Info("AddList")

	for _, v := range applications {
		_, err := svc.repo.Get(cxt, v.UUID, true)

		switch {
		case err == nil:
			_, err = svc.repo.Update(cxt, v)
			if err != nil {
				svc.log.WithError(err).WithField("uuid", v.UUID.String()).Error("AddList: error while updating")
				return err
			}

		case err == app_errors.ErrNotFound:
			_, err = svc.repo.Create(cxt, v)
			if err != nil {
				svc.log.WithError(err).WithField("uuid", v.UUID.String()).Error("AddList: error while creating")
				return err
			}

		default:
			svc.log.WithError(err).WithField("uuid", v.UUID.String()).Error("AddList: error")
			return err
		}
	}

	return nil
}

func (svc *ApplicationSvc) Delete(cxt context.Context, uuid uuid.UUID) error {
	svc.log.WithField("uuid", uuid.String()).Info("Delete")

	err := svc.repo.Delete(cxt, uuid)
	if err == app_errors.ErrNotFound {
		svc.log.WithField("uuid", uuid.String()).Info("Delete: application not found")
	} else if err != nil {
		svc.log.WithError(err).WithField("uuid", uuid.String()).Error("Delete: error")
	}

	return err
}
