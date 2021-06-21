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

	output := map[string]map[string]map[string][]string{}
	output["data"] = map[string]map[string][]string{}
	for i, credentials := range data.Credentials {
		key := credentials.Url + ":" + credentials.CommitSHA
		val := map[string][]string{}
		for _, file := range data.Repositories[i].Files {
			val[file.Path] = file.Objects
		}
		output["data"][key] = val
	}

	j, err := json.MarshalIndent(output, "", "\t")
	data.Output = string(j)
	return err
}
