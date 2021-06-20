package readers

import (
	"net/http"
	"redhat-sre-task-dockerfile-scanner/src/models"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Reader interface {
	Read(data *models.Data) error
}
