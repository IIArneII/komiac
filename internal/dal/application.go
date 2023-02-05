package dal

import (
	"context"
	"errors"
	"komiac/internal/app/entities"

	uuid "github.com/satori/go.uuid"
)

func (s *Storage) GetList(cxt context.Context, filter entities.ApplicationFilter) ([]*entities.Application, error) {
	return nil, errors.New("not implemented")
}

func (s *Storage) Create(cxt context.Context, application entities.Application) (*entities.Application, error) {
	return nil, errors.New("not implemented")
}

func (s *Storage) Get(cxt context.Context, uuid uuid.UUID) (*entities.Application, error) {
	return nil, errors.New("not implemented")
}

func (s *Storage) Delete(cxt context.Context, uuid uuid.UUID) error {
	return errors.New("not implemented")
}

func (s *Storage) Update(cxt context.Context, application entities.Application) (*entities.Application, error) {
	return nil, errors.New("not implemented")
}
