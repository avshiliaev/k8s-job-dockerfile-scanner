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
	}
	reader := readers.RemoteTxtReader(&ClientMock{})

	// Act
	err := reader.Read(data)

	// Assert
	if len(data.InputLines) != 3 || err != nil {
		t.Error()
	}
}
