package test

import (
	"redhat-sre-task-dockerfile-scanner/src/models"
	"redhat-sre-task-dockerfile-scanner/src/parsers"
	"testing"
)

var dockerFileBody = `
FROM alpine
ADD scanner-job /scanner-job
ENTRYPOINT [ "/scanner-job" ]
`

func TestDockerFileParser(t *testing.T) {

	// Arrange
	data := &models.Data{
		Repositories: []models.Repo{
			{
				Files: []models.File{
					{
						Content: dockerFileBody,
					},
				},
			},
		},
	}
	parser := parsers.DockerFileParser()

	// Act
	parser.Parse(data)

	// Assert
	for _, repo := range data.Repositories {
		for _, file := range repo.Files {
			if len(file.Objects) == 0 {
				t.Error()
			}
		}
	}

}
