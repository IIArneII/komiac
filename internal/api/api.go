package api

import (
	"komiac/internal/api/restapi/restapi"
	"komiac/internal/api/restapi/restapi/operations"
	"komiac/internal/api/restapi/restapi/operations/application"
	"komiac/internal/app"

	"github.com/go-openapi/loads"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host string
	Port int
}

type service struct {
	log *logrus.Logger
	app *app.App
}

func NewServer(app *app.App, cfg Config) (*restapi.Server, error) {
	service := &service{
		app: app,
		log: logrus.New(),
	}

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load embedded swagger spec")
	}

	api := operations.NewKomiacAPI(swaggerSpec)

	api.ApplicationAddListHandler = application.AddListHandlerFunc(service.ApplicationAddList)
	api.ApplicationGetListHandler = application.GetListHandlerFunc(service.ApplicationGetList)
	api.ApplicationDeleteHandler = application.DeleteHandlerFunc(service.ApplicationDelete)

	api.Logger = func(s string, i ...interface{}) {
		service.log.Info(append(append(make([]interface{}, 0), s), i...)...)
	}

	server := restapi.NewServer(api)
	server.Host = cfg.Host
	server.Port = cfg.Port

	return server, nil
}
