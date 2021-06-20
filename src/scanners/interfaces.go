package scanners

import (
	"redhat-sre-task-dockerfile-scanner/src/api"
	"redhat-sre-task-dockerfile-scanner/src/parsers"
	"redhat-sre-task-dockerfile-scanner/src/readers"
	"redhat-sre-task-dockerfile-scanner/src/validators"
	"redhat-sre-task-dockerfile-scanner/src/writers"
)

// Scanner abstract type
type Scanner interface {
	Read(reader readers.Reader) error
	Validate(validator validators.Validator) error
	Query(api api.RepositoryApi) error
	Parse(parser parsers.FileParser) error
	Write(writer writers.Writer) error
}
