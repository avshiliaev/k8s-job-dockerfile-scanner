package github

import (
	"context"
	"github.com/google/go-github/v35/github"
	"redhat-sre-task-dockerfile-scanner/src/models"
	"strings"
)

type AbstractGitHubClient interface {
	CheckFileFormat(repo models.RepoCredentials, namePattern string) (bool, error)
	GetFilePaths(repo models.RepoCredentials, namePattern string) ([]string, error)
	GetContent(repo models.RepoCredentials, path string) (string, error)
}

type googleGitHubClient struct {
	Client *github.Client
}

func GoogleGitHubClient() *googleGitHubClient {
	return &googleGitHubClient{
		Client: github.NewClient(nil),
	}
}

func (api *googleGitHubClient) CheckFileFormat(repo models.RepoCredentials, namePattern string) (bool, error) {
	languages, _, err := api.Client.Repositories.ListLanguages(context.Background(), repo.Owner, repo.Name)
	if _, ok := languages[namePattern]; ok {
		return true, err
	}
	return false, err
}

func (api *googleGitHubClient) GetFilePaths(repo models.RepoCredentials, namePattern string) ([]string, error) {
	tree, _, err := api.Client.Git.GetTree(
		context.Background(),
		repo.Owner,
		repo.Name,
		repo.CommitSHA,
		true,
	)
	if err != nil {
		return nil, err
	}
	var paths []string
	for _, entry := range tree.Entries {
		path := entry.GetPath()
		pathParsed := strings.Split(path, "/")
		if pathParsed[len(pathParsed)-1] == namePattern {
			paths = append(paths, path)
		}
	}

	return paths, err
}

func (api *googleGitHubClient) GetContent(repo models.RepoCredentials, path string) (string, error) {
	opt := &github.RepositoryContentGetOptions{Ref: repo.CommitSHA}
	contentEncoded, _, _, err := api.Client.Repositories.GetContents(
		context.Background(),
		repo.Owner,
		repo.Name,
		path,
		opt,
	)
	if err != nil {
		return "", err
	}
	content, err := contentEncoded.GetContent()
	if err != nil {
		return "", err
	}

	return content, err
}
