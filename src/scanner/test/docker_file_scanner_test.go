package test

import (
	"redhat-sre-task-dockerfile-scanner/src/api/github"
	"redhat-sre-task-dockerfile-scanner/src/readers"
	scanners "redhat-sre-task-dockerfile-scanner/src/scanner"
	"redhat-sre-task-dockerfile-scanner/src/validators"
	"testing"
)

func TestScanner(t *testing.T) {

	// Arrange
	scanner := scanners.DockerFileScanner("https://test.com/test.txt")

	// Act
	//TODO: add chain of calls
	scanner.Read(readers.RemoteTxtReader(&MockHttpClient{}))
	scanner.Validate(validators.GitHubValidator())
	scanner.Query(github.Api(&MockGitHubClient{}))

	// Assert
	if scanner.GetData().Output == "" {
		t.Error()
	}
}
