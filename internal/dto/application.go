package dto

import (
	"komiac/internal/api/restapi/models"
	"komiac/internal/app/entities"
)

func ApiApplication(application entities.Application) models.Application {
	return models.Application{}
}

func AppApplication(application models.Application) entities.Application {
	return entities.Application{}
}

func ApiApplications(application []entities.Application) *models.Applications {
	return &models.Applications{}
}

func AppApplications(application models.Applications) []entities.Application {
	return []entities.Application{}
}
