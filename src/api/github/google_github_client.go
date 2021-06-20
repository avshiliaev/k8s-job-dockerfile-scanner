package github

import (
	"context"
	"fmt"
	"github.com/google/go-github/v35/github"
	"redhat-sre-task-dockerfile-scanner/src/models"
	"strings"
)

type AbstractGitHubClient interface {
	CheckFileFormat(repo models.RepoCredentials, namePattern string) bool
	GetFilePaths(repo models.RepoCredentials, namePattern string) []string
	GetContent(repo models.RepoCredentials, path string) string
}

type googleGitHubClient struct {
	Client *github.Client
}

func GoogleGitHubClient() *googleGitHubClient {
	return &googleGitHubClient{
		Client: github.NewClient(nil),
	}
}

func (api *googleGitHubClient) CheckFileFormat(repo models.RepoCredentials, namePattern string) bool {
	languages, response, err := api.Client.Repositories.ListLanguages(context.Background(), repo.Owner, repo.Name)
	fmt.Println(languages, response, err)
	if _, ok := languages[namePattern]; ok {
		return true
	}
	return false
}

func (api *googleGitHubClient) GetFilePaths(repo models.RepoCredentials, namePattern string) []string {
	tree, response, err := api.Client.Git.GetTree(
		context.Background(),
		repo.Owner,
		repo.Name,
		repo.CommitSHA,
		true,
	)
	if err != nil {
		fmt.Println(response, err)
	}
	var paths []string
	for _, entry := range tree.Entries {
		path := entry.GetPath()
		pathParsed := strings.Split(path, "/")
		if pathParsed[len(pathParsed)-1] == namePattern {
			paths = append(paths, path)
		}
	}

	return paths
}

func (api *googleGitHubClient) GetContent(repo models.RepoCredentials, path string) string {
	opt := &github.RepositoryContentGetOptions{Ref: repo.CommitSHA}
	contentEncoded, _, response, err := api.Client.Repositories.GetContents(
		context.Background(),
		repo.Owner,
		repo.Name,
		path,
		opt,
	)
	if err != nil {
		fmt.Println(response, err)
	}
	content, err := contentEncoded.GetContent()
	if err != nil {
		fmt.Println(response, err)
	}

	return content
}
