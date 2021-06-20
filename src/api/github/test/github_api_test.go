package test

import (
	"redhat-sre-task-dockerfile-scanner/src/api/github"
	"redhat-sre-task-dockerfile-scanner/src/models"
	"testing"
)

func TestApi(t *testing.T) {

	// Arrange
	data := &models.Data{
		Credentials: []models.RepoCredentials{
			{
				Owner:     "awesome",
				Name:      "thing",
				CommitSHA: "123",
			},
		},
	}
	api := github.Api(&MockGitHubClient{})

	// Act
	err := api.Query(data)

	// Assert
	if len(data.Repositories) == 0 || err != nil {
		t.Error()
	}

}
