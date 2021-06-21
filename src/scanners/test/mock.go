package test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"redhat-sre-task-dockerfile-scanner/src/models"
)

type MockHttpClient struct {
}

func (c *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	var responseBody = `
https://github.com/owner/repo1 3425346356456
https://github.com/owner/repo2 3425346356456
`
	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(responseBody)),
	}, nil
}

type MockGitHubClient struct {
}

func (api *MockGitHubClient) CheckFileFormat(repo models.RepoCredentials, namePattern string) (bool, error) {
	return true, nil
}

func (api *MockGitHubClient) GetFilePaths(repo models.RepoCredentials, namePattern string) ([]string, error) {
	return []string{"awesome/Dockerfile"}, nil
}

func (api *MockGitHubClient) GetContent(repo models.RepoCredentials, path string) (string, error) {
	var dockerFileBody = `
FROM alpine
ADD scanner-job /scanner-job
COPY ./input.txt /input.txt
ENTRYPOINT [ "/scanner-job" ]
`
	return dockerFileBody, nil
}
