package api

import "redhat-sre-task-dockerfile-scanner/src/models"

type RepositoryApi interface {
	Query(data *models.Data)
}
