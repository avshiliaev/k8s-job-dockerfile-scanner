package parsers

import (
	"redhat-sre-task-dockerfile-scanner/src/models"
	"redhat-sre-task-dockerfile-scanner/src/parsers/buildkit"
	"strings"
)

type dockerFileParser struct {
}

func DockerFileParser() *dockerFileParser {
	return &dockerFileParser{}
}

func (p *dockerFileParser) Parse(data *models.Data) error {

	var err error
	var repos []models.Repo
	for _, repo := range data.Repositories {
		var files []models.File
		for _, file := range repo.Files {
			parsed, _ := buildkit.Parse(strings.NewReader(file.Content))
			var images []string
			for _, child := range parsed.AST.Children {
				if child.Value == "from" {
					images = append(images, child.Next.Value)
				}
			}
			files = append(files, models.File{
				Path:    file.Path,
				Objects: images,
			})
		}
		repos = append(repos, models.Repo{
			Files: files,
		})
	}
	data.Repositories = repos
	return err
}
