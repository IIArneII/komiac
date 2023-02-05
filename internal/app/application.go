package app

import (
	"context"
	"errors"
	"komiac/internal/app/entities"

	uuid "github.com/satori/go.uuid"
)

type ApplicationService interface {
	GetList(cxt context.Context, filter entities.ApplicationFilter) (*entities.Application, error)
	AddList(cxt context.Context, application []entities.Application) (*entities.Application, error)
	Delete(cxt context.Context, uuid uuid.UUID) (*entities.Application, error)
}

type ApplicationRepository interface {
	GetList(cxt context.Context, filter entities.ApplicationFilter) ([]*entities.Application, error)
	Get(cxt context.Context, uuid uuid.UUID) (*entities.Application, error)
	Create(cxt context.Context, application entities.Application) (*entities.Application, error)
	Delete(cxt context.Context, uuid uuid.UUID) error
	Update(cxt context.Context, application entities.Application) (*entities.Application, error)
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

func (svc *ApplicationSvc) AddList(cxt context.Context, application []entities.Application) (*entities.Application, error) {
	return nil, errors.New("not implemented")
}

func (svc *ApplicationSvc) Delete(cxt context.Context, uuid uuid.UUID) (*entities.Application, error) {
	return nil, errors.New("not implemented")
}
