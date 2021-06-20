package github

import (
	"redhat-sre-task-dockerfile-scanner/src/models"
)

type AbstractGitHubClient interface {
	CheckFileFormat(repo models.RepoCredentials, namePattern string) bool
	GetFilePaths(repo models.RepoCredentials, namePattern string) []string
	GetContent(repo models.RepoCredentials, path string) string
}
