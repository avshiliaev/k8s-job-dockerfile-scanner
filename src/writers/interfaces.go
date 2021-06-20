package writers

import "redhat-sre-task-dockerfile-scanner/src/models"

type Writer interface {
	Convert(data *models.Data)
}
