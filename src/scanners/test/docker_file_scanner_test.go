package test

import (
	"redhat-sre-task-dockerfile-scanner/src/api/github"
	"redhat-sre-task-dockerfile-scanner/src/parsers"
	"redhat-sre-task-dockerfile-scanner/src/readers"
	scanners "redhat-sre-task-dockerfile-scanner/src/scanners"
	"redhat-sre-task-dockerfile-scanner/src/validators"
	"redhat-sre-task-dockerfile-scanner/src/writers"
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
	scanner.Parse(parsers.DockerFileParser())
	scanner.Write(writers.JsonStdWriter())

	// Assert
	if scanner.GetData().Output == "" {
		t.Error()
	}
}
