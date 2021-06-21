package serializers

import "redhat-sre-task-dockerfile-scanner/src/models"

type Serializer interface {
	Serialize(data *models.Data) error
}
