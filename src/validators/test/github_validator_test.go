package test

import (
	"redhat-sre-task-dockerfile-scanner/src/models"
	"redhat-sre-task-dockerfile-scanner/src/validators"
	"testing"
)

func TestGitHubValidator(t *testing.T) {

	// Arrange
	data := &models.Data{
		InputLines: []string{
			"https://github.com/owner/repo1 3425346356456",
			"https://github.com/owner/repo2 3425346356456",
		},
	}
	validator := validators.GitHubValidator()

	// Act
	validator.Validate(data)

	// Assert
	if len(data.Credentials) == 0 {
		t.Error()
	}
}
