package api

import (
	"komiac/internal/api/restapi/models"
	"komiac/internal/api/restapi/restapi/operations/application"
	"komiac/internal/app/entities"
	app_errors "komiac/internal/app/errors"
	"komiac/internal/dto"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func (svc *service) ApplicationAddList(params application.AddListParams) application.AddListResponder {
	fields := logrus.Fields{
		"method": params.HTTPRequest.Method,
		"url":    params.HTTPRequest.URL,
	}

	applications, err := dto.ApiToAppApplications(params.Body)
	if err != nil {
		fields["code"] = 400
		svc.log.WithError(err).WithFields(fields).Error("AddList")
		errStr := err.Error()
		return application.NewAddListDefault(500).WithPayload(&models.Error{
			Message: &errStr,
		})
	}

	err = svc.app.ApplicationSvc.AddList(params.HTTPRequest.Context(), applications)

	if err != nil {
		fields["code"] = 500
		svc.log.WithError(err).WithFields(fields).Error("AddList")
		errStr := err.Error()
		return application.NewAddListDefault(500).WithPayload(&models.Error{
			Message: &errStr,
		})
	}

	fields["code"] = 200
	svc.log.WithFields(fields).Info("AddList")
	return application.NewAddListOK()
}

func (svc *service) ApplicationGetList(params application.GetListParams) application.GetListResponder {
	fields := logrus.Fields{
		"method": params.HTTPRequest.Method,
		"url":    params.HTTPRequest.URL,
	}

	applications, err := svc.app.ApplicationSvc.GetList(params.HTTPRequest.Context(), entities.ApplicationFilter{
		DivisionOID: params.DivisionOID,
		MNN:         params.MNN,
		Year:        params.Year,
	})
	if err != nil {
		fields["code"] = 500
		svc.log.WithError(err).WithFields(fields).Error("GetList")
		errStr := err.Error()
		return application.NewGetListDefault(500).WithPayload(&models.Error{
			Message: &errStr,
		})
	}

	fields["code"] = 200
	svc.log.WithFields(fields).Info("GetList")
	return application.NewGetListOK().WithPayload(dto.AppToApiApplications(applications))
}

func (svc *service) ApplicationDelete(params application.DeleteParams) application.DeleteResponder {
	fields := logrus.Fields{
		"method": params.HTTPRequest.Method,
		"url":    params.HTTPRequest.URL,
	}

	uuid, err := uuid.FromString(params.UUID.String())
	if err != nil {
		fields["code"] = 400
		svc.log.WithError(err).WithFields(fields).Error("Delete")
		return application.NewDeleteDefault(400)
	}

	err = svc.app.ApplicationSvc.Delete(params.HTTPRequest.Context(), uuid)

	if err == app_errors.ErrNotFound {
		fields["code"] = 404
		svc.log.WithFields(fields).Info("Delete")
		errStr := err.Error()
		return application.NewDeleteNotFound().WithPayload(&models.Error{
			Message: &errStr,
		})
	}

	if err != nil {
		fields["code"] = 500
		svc.log.WithError(err).WithFields(fields).Error("Delete")
		errStr := err.Error()
		return application.NewDeleteDefault(500).WithPayload(&models.Error{
			Message: &errStr,
		})
	}

	fields["code"] = 200
	svc.log.WithFields(fields).Info("Delete")
	return application.NewDeleteOK()
}
