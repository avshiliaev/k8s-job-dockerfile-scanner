package github

import "redhat-sre-task-dockerfile-scanner/src/models"

type gitHubApi struct {
	Client AbstractGitHubClient
}

func Api(client AbstractGitHubClient) *gitHubApi {
	return &gitHubApi{
		Client: client,
	}
}

func (api *gitHubApi) Query(data *models.Data) {

	var repos []models.Repo
	for _, credentials := range data.Credentials {
		files := api.fetchFiles(credentials)
		repos = append(repos, models.Repo{
			Files: files,
		})
	}
	data.Repositories = repos
}

func (api *gitHubApi) fetchFiles(credentials models.RepoCredentials) []models.File {
	var files []models.File
	fileNamePattern := "Dockerfile"
	dockerExists := api.Client.CheckFileFormat(credentials, fileNamePattern)

	if dockerExists {
		paths := api.Client.GetFilePaths(credentials, fileNamePattern)

		for _, path := range paths {
			content := api.Client.GetContent(credentials, path)
			files = append(files, models.File{Path: path, Content: content})
		}
	}
	return files
}
