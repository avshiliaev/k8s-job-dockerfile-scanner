package test

import (
	"redhat-sre-task-dockerfile-scanner/src/models"
	"redhat-sre-task-dockerfile-scanner/src/serializers"
	"testing"
)

func TestJsonStdWriter(t *testing.T) {

	// Arrange
	data := &models.Data{
		Credentials: []models.RepoCredentials{
			{
				Url:       "https://github.com/awesome/project",
				Owner:     "awesome",
				Name:      "project",
				CommitSHA: "123",
			},
		},
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
	output := `{
	"data": {
		"https://github.com/awesome/project:123": {
			"path/to/file": [
				"object"
			]
		}
	}
}`
	serializer := serializers.JsonSerializers()

	// Act
	err := serializer.Serialize(data)

	// Assert
	if data.Output != output || err != nil {
		t.Error()
	}
}
