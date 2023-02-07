package app

import (
	"context"
	"errors"
	"komiac/internal/app/entities"
	app_errors "komiac/internal/app/errors"

	uuid "github.com/satori/go.uuid"
)

type ApplicationService interface {
	GetList(cxt context.Context, filter entities.ApplicationFilter) (*entities.Application, error)
	AddList(cxt context.Context, application []*entities.Application) error
	Delete(cxt context.Context, uuid uuid.UUID) (*entities.Application, error)
}

type ApplicationRepository interface {
	GetList(cxt context.Context, filter entities.ApplicationFilter) ([]*entities.Application, error)
	Get(cxt context.Context, uuid uuid.UUID) (*entities.Application, error)
	Create(cxt context.Context, application *entities.Application) (*entities.Application, error)
	Delete(cxt context.Context, uuid uuid.UUID) error
	Update(cxt context.Context, application *entities.Application) (*entities.Application, error)
}

type ApplicationSvc struct {
	repo ApplicationRepository
}

func NewApplicationService(repo ApplicationRepository) ApplicationService {
	return &ApplicationSvc{
		repo: repo,
	}
}

func (svc *ApplicationSvc) GetList(cxt context.Context, filter entities.ApplicationFilter) (*entities.Application, error) {
	return nil, errors.New("not implemented")
}

func (svc *ApplicationSvc) AddList(cxt context.Context, applications []*entities.Application) error {
	for _, v := range applications {
		_, err := svc.repo.Get(cxt, v.UUID)

		switch {
		case err == nil:
			_, err = svc.repo.Update(cxt, v)
			if err != nil {
				return err
			}

		case err == app_errors.ErrNotFound:
			_, err = svc.repo.Create(cxt, v)
			if err != nil {
				return err
			}

		default:
			return err
		}
	}

	return nil
}

func (svc *ApplicationSvc) Delete(cxt context.Context, uuid uuid.UUID) (*entities.Application, error) {
	return nil, errors.New("not implemented")
}
