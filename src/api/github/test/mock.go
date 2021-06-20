package test

import "redhat-sre-task-dockerfile-scanner/src/models"

type MockGitHubClient struct {
}

func (api *MockGitHubClient) CheckFileFormat(repo models.RepoCredentials, namePattern string) bool {
	return true
}

func (api *MockGitHubClient) GetFilePaths(repo models.RepoCredentials, namePattern string) []string {
	return []string{"awesome/Dockerfile"}
}

func (api *MockGitHubClient) GetContent(repo models.RepoCredentials, path string) string {
	var dockerFileBody = `
FROM alpine
ADD scanner-job /scanner-job
COPY ./input.txt /input.txt
ENTRYPOINT [ "/scanner-job" ]
`
	return dockerFileBody
}
