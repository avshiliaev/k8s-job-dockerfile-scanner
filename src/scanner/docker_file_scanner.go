package scanners

import (
	"redhat-sre-task-dockerfile-scanner/src/api"
	"redhat-sre-task-dockerfile-scanner/src/models"
	"redhat-sre-task-dockerfile-scanner/src/parsers"
	"redhat-sre-task-dockerfile-scanner/src/readers"
	"redhat-sre-task-dockerfile-scanner/src/validators"
	"redhat-sre-task-dockerfile-scanner/src/writers"
)

// dockerFileScanner Concrete private implementation
type dockerFileScanner struct {
	data *models.Data
}

// DockerFileScanner Constructor Functions for a concrete implementation
func DockerFileScanner(url string) *dockerFileScanner {
	data := &models.Data{Url: url}
	return &dockerFileScanner{data: data}
}

// Methods with pointer receivers
func (s *dockerFileScanner) Read(reader readers.Reader) {
	reader.Read(s.data)
}
func (s *dockerFileScanner) Validate(validator validators.Validator) {
	validator.Validate(s.data)
}
func (s *dockerFileScanner) Query(api api.RepositoryApi) {
	api.Query(s.data)
}
func (s *dockerFileScanner) Parse(parser parsers.FileParser) {
	parser.Parse(s.data)
}
func (s *dockerFileScanner) Write(writer writers.Writer) {
	writer.Convert(s.data)
}
func (s *dockerFileScanner) GetData() *models.Data {
	return s.data
}
