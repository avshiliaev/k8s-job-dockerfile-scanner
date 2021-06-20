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
func (s *dockerFileScanner) Read(reader readers.Reader) error {
	err := reader.Read(s.data)
	return err
}
func (s *dockerFileScanner) Validate(validator validators.Validator) error {
	err := validator.Validate(s.data)
	return err
}
func (s *dockerFileScanner) Query(api api.RepositoryApi) error {
	err := api.Query(s.data)
	return err
}
func (s *dockerFileScanner) Parse(parser parsers.FileParser) error {
	err := parser.Parse(s.data)
	return err
}
func (s *dockerFileScanner) Write(writer writers.Writer) error {
	err := writer.Write(s.data)
	return err
}
func (s *dockerFileScanner) GetData() *models.Data {
	return s.data
}
