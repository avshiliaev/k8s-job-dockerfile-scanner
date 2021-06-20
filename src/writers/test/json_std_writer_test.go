package test

import (
	"redhat-sre-task-dockerfile-scanner/src/models"
	"redhat-sre-task-dockerfile-scanner/src/writers"
	"testing"
)

func TestJsonStdWriter(t *testing.T) {

	// Arrange
	data := &models.Data{
		Url:          "https://test.com/",
		InputLines:   nil,
		Credentials:  nil,
		Repositories: nil,
		Output:       "",
	}
	writer := writers.JsonStdWriter()

	// Act
	writer.Write(data)

	// Assert
	if data.Output == "" {
		t.Error()
	}
}
