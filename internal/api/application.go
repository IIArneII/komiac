package api

import (
	"komiac/internal/api/restapi/models"
	"komiac/internal/api/restapi/restapi/operations/application"
	"komiac/internal/app/entities"
	"komiac/internal/dto"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func (svc *service) ApplicationAddList(params application.AddListParams) application.AddListResponder {
	svc.log.WithFields(logrus.Fields{
		"method": params.HTTPRequest.Method,
		"url":    params.HTTPRequest.URL,
	}).Info("ApplicationAddList")

	applications, err := dto.ApiToAppApplications(params.Body)
	if err != nil {
		errStr := err.Error()
		return application.NewAddListDefault(500).WithPayload(&models.Error{
			Message: &errStr,
		})
	}

	err = svc.app.ApplicationSvc.AddList(params.HTTPRequest.Context(), applications)
	if err != nil {
		errStr := err.Error()
		return application.NewAddListDefault(500).WithPayload(&models.Error{
			Message: &errStr,
		})
	}

	return application.NewAddListOK()
}

func (svc *service) ApplicationGetList(params application.GetListParams) application.GetListResponder {
	svc.log.WithFields(logrus.Fields{
		"method": params.HTTPRequest.Method,
		"url":    params.HTTPRequest.URL,
	}).Info("ApplicationGetList")

	applications, err := svc.app.ApplicationSvc.GetList(params.HTTPRequest.Context(), entities.ApplicationFilter{
		DivisionOID: params.DivisionOID,
		MNN:         params.MNN,
		Year:        params.Year,
	})
	if err != nil {
		errStr := err.Error()
		return application.NewGetListDefault(500).WithPayload(&models.Error{
			Message: &errStr,
		})
	}

	return application.NewGetListOK().WithPayload(dto.AppToApiApplications(applications))
}

func (svc *service) ApplicationDelete(params application.DeleteParams) application.DeleteResponder {
	svc.log.WithFields(logrus.Fields{
		"method": params.HTTPRequest.Method,
		"url":    params.HTTPRequest.URL,
	}).Info("ApplicationDelete")

	uuid, err := uuid.FromString(params.UUID.String())
	if err != nil {
		return application.NewDeleteDefault(400)
	}

	err = svc.app.ApplicationSvc.Delete(params.HTTPRequest.Context(), uuid)
	if err != nil {
		errStr := err.Error()
		return application.NewDeleteDefault(500).WithPayload(&models.Error{
			Message: &errStr,
		})
	}

	return application.NewDeleteOK()
}
