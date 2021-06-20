package test

import (
	"redhat-sre-task-dockerfile-scanner/src/models"
	"redhat-sre-task-dockerfile-scanner/src/serializers"
	"testing"
)

func TestJsonStdWriter(t *testing.T) {

	// Arrange
	data := &models.Data{
		Repositories: []models.Repo{
			{
				Files: []models.File{
					{
						Path:    "path/to/file",
						Content: "",
						Objects: []string{"object"},
					},
				},
			},
		},
	}
	serializer := serializers.JsonSerializers()

	// Act
	err := serializer.Serialize(data)

	// Assert
	if data.Output == "" || err != nil {
		t.Error()
	}
}
