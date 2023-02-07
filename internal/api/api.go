package api

import (
	"komiac/internal/api/restapi/restapi"
	"komiac/internal/api/restapi/restapi/operations"
	"komiac/internal/api/restapi/restapi/operations/application"
	"komiac/internal/app"

	"github.com/go-openapi/loads"
	"github.com/pkg/errors"
)

type Config struct {
	Host string
	Port int
}

type service struct {
	app *app.App
}

func NewServer(app *app.App, cfg Config) (*restapi.Server, error) {
	service := &service{
		app: app,
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load embedded swagger spec")
	}

	api := operations.NewKomiacAPI(swaggerSpec)

	api.ApplicationAddListHandler = application.AddListHandlerFunc(service.ApplicationAddList)
	api.ApplicationGetListHandler = application.GetListHandlerFunc(service.ApplicationGetList)
	api.ApplicationDeleteHandler = application.DeleteHandlerFunc(service.ApplicationDelete)

	server := restapi.NewServer(api)
	server.Host = cfg.Host
	server.Port = cfg.Port

	return server, nil
}
