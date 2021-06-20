package test

import (
	scanners "redhat-sre-task-dockerfile-scanner/src/scanner"
	"testing"
)

func TestScanner(t *testing.T) {

	// Arrange
	scanner := scanners.DockerFileScanner("https://test.com/test.txt")

	// Act
	//TODO: add chain of calls

	// Assert
	if scanner.GetData().Output == "" {
		t.Error()
	}
}
