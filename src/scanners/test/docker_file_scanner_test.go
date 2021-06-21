package test

import (
	"redhat-sre-task-dockerfile-scanner/src/api/github"
	"redhat-sre-task-dockerfile-scanner/src/parsers"
	"redhat-sre-task-dockerfile-scanner/src/readers"
	"redhat-sre-task-dockerfile-scanner/src/scanners"
	"redhat-sre-task-dockerfile-scanner/src/serializers"
	"redhat-sre-task-dockerfile-scanner/src/validators"
	"testing"
)

func TestScanner(t *testing.T) {

	// Arrange
	scanner := scanners.DockerFileScanner("https://test.com/test.txt")

	// Act
	var err error
	err = scanner.Read(readers.RemoteTxtReader(&MockHttpClient{}))
	err = scanner.Validate(validators.GitHubValidator())
	err = scanner.Query(github.Api(&MockGitHubClient{}))
	err = scanner.Parse(parsers.DockerFileParser())
	err = scanner.Serialize(serializers.JsonSerializers())

	// Assert
	if scanner.GetData().Output == "" || err != nil {
		t.Error()
	}
}
