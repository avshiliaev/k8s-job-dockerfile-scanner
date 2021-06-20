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

func (api *gitHubApi) Query(data *models.Data) error {

	var err error
	var repos []models.Repo
	for _, credentials := range data.Credentials {
		files, err := api.fetchFiles(credentials)
		if err != nil {
			return err
		}
		repos = append(repos, models.Repo{
			Files: files,
		})
	}
	data.Repositories = repos
	return err
}

func (api *gitHubApi) fetchFiles(credentials models.RepoCredentials) ([]models.File, error) {

	var err error
	var files []models.File
	fileNamePattern := "Dockerfile"
	dockerExists, errLocal := api.Client.CheckFileFormat(credentials, fileNamePattern)
	err = errLocal

	if dockerExists {
		paths, errLocal := api.Client.GetFilePaths(credentials, fileNamePattern)
		err = errLocal

		for _, path := range paths {
			content, errLocal := api.Client.GetContent(credentials, path)
			err = errLocal
			files = append(files, models.File{Path: path, Content: content})
		}
	}
	return files, err
}
