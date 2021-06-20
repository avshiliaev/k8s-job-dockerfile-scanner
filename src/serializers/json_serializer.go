package serializers

import (
	"encoding/json"
	"redhat-sre-task-dockerfile-scanner/src/models"
)

type jsonSerializers struct {
}

func JsonSerializers() *jsonSerializers {
	return &jsonSerializers{}
}

func (jsonSerializers *jsonSerializers) Serialize(data *models.Data) error {
	j, err := json.MarshalIndent(data, "", "\t")
	data.Output = string(j)
	return err
}
