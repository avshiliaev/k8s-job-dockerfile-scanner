package test

import (
	"redhat-sre-task-dockerfile-scanner/src/models"
	"redhat-sre-task-dockerfile-scanner/src/readers"
	"testing"
)

func TestRemoteTxtReader(t *testing.T) {

	// Arrange
	data := &models.Data{
		Url:          "https://test.com/",
		InputLines:   nil,
		Credentials:  nil,
		Repositories: nil,
		Output:       "",
	}
	reader := readers.RemoteTxtReader(&ClientMock{})

	// Act
	reader.Read(data)

	// Assert
	if len(data.InputLines) != 3 {
		t.Error()
	}
}
