package parsers

import "redhat-sre-task-dockerfile-scanner/src/models"

type FileParser interface {
	Parse(data *models.Data) error
}
