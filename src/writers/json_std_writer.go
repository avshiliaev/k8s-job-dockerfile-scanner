package writers

import (
	"encoding/json"
	"redhat-sre-task-dockerfile-scanner/src/models"
)

type jsonStdWriter struct {
}

func JsonStdWriter() *jsonStdWriter {
	return &jsonStdWriter{}
}

func (jsonWriter *jsonStdWriter) Write(data *models.Data) error {
	j, err := json.MarshalIndent(data, "", "\t")
	data.Output = string(j)
	return err
}
