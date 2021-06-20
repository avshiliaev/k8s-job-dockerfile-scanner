package writers

import "redhat-sre-task-dockerfile-scanner/src/models"

type Writer interface {
	Write(data *models.Data) error
}
