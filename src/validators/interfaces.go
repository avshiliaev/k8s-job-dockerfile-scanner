package validators

import "redhat-sre-task-dockerfile-scanner/src/models"

type Validator interface {
	Validate(data *models.Data)
}
