package validators

import (
	"net/url"
	"path/filepath"
	"redhat-sre-task-dockerfile-scanner/src/models"
	"strings"
)

type gitHubValidator struct {
}

func GitHubValidator() *gitHubValidator {
	return &gitHubValidator{}
}

func (validator *gitHubValidator) Validate(data *models.Data) error {

	var err error
	var credentials []models.RepoCredentials
	for _, inputLine := range data.InputLines {

		if isEmptyString(inputLine) {
			continue
		}
		repoURL, commitSHA := splitOnCharacter(inputLine, " ")
		if isEmptyString(repoURL) || isEmptyString(commitSHA) {
			continue
		}
		owner, name, errLocal := getOwnerAndName(repoURL)
		err = errLocal

		credentials = append(credentials, models.RepoCredentials{
			Url:       repoURL,
			Owner:     owner,
			Name:      name,
			CommitSHA: commitSHA,
		})
	}
	data.Credentials = credentials

	return err

}

func isEmptyString(s string) bool {
	if s == "" {
		return true
	}
	return false
}

func splitOnCharacter(s string, char string) (string, string) {
	split := strings.Split(s, char)
	left, right := split[0], split[1]
	return left, right
}

func getOwnerAndName(repoURL string) (string, string, error) {
	repoParsedURL, err := url.Parse(repoURL)
	if err != nil {
		return "", "", err
	}

	parts := strings.Split(repoParsedURL.Path, "/")
	owner := parts[1]
	name := parts[2]

	nameTrimmed := strings.TrimSuffix(name, filepath.Ext(name))

	return owner, nameTrimmed, nil
}
