package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	var responseBody = `
https://github.com/owner/repo1 3425346356456
https://github.com/owner/repo2 3425346356456
`
	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(responseBody)),
	}, nil
}
