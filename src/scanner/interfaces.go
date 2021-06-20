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
	Read(reader readers.Reader)
	Validate(validator validators.Validator)
	Query(api api.RepositoryApi)
	Parse(parser parsers.FileParser)
	Write(writer writers.Writer)
}
