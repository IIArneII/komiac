package api

import (
	"komiac/internal/api/restapi/models"
	"komiac/internal/api/restapi/restapi/operations/application"
	"komiac/internal/app/entities"
	"komiac/internal/dto"

	uuid "github.com/satori/go.uuid"
)

func (svc *service) ApplicationAddList(params application.AddListParams) application.AddListResponder {

	applications, err := dto.AppApplications(params.Body)
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
	svc.app.ApplicationSvc.GetList(params.HTTPRequest.Context(), entities.ApplicationFilter{
		DivisionOID: params.DivisionOID,
		MNN:         params.MNN,
		Year:        params.Year,
	})

	return application.NewGetListDefault(404)
}

func (svc *service) ApplicationDelete(params application.DeleteParams) application.DeleteResponder {
	uuid, err := uuid.FromString(params.UUID.String())
	if err != nil {
		return application.NewDeleteDefault(400)
	}

	svc.app.ApplicationSvc.Delete(params.HTTPRequest.Context(), uuid)

	return application.NewDeleteDefault(404)
}
