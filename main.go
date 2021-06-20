package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"redhat-sre-task-dockerfile-scanner/src/api/github"
	"redhat-sre-task-dockerfile-scanner/src/parsers"
	"redhat-sre-task-dockerfile-scanner/src/readers"
	"redhat-sre-task-dockerfile-scanner/src/scanners"
	"redhat-sre-task-dockerfile-scanner/src/serializers"
	"redhat-sre-task-dockerfile-scanner/src/validators"
)

const (
	input   = "input"
	format  = "format"
	pattern = "pattern"
	vendor  = "vendor"
	out     = "out"
)

func main() {

	// ./scanner -i http://github.com/ -p Dockerfile -v github -o json
	app := &cli.App{
		Name:  "Scanner",
		Usage: "scans repositories for Dockerfiles and retrieves image information",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     input,
				Aliases:  []string{"i"},
				Value:    "",
				Usage:    "link to a txt input file",
				Required: true,
			},
			&cli.StringFlag{
				Name:    format,
				Aliases: []string{"f"},
				Value:   "txt",
				Usage:   "input format",
			},
			&cli.StringFlag{
				Name:    pattern,
				Aliases: []string{"p"},
				Value:   "Dockerfile",
				Usage:   "a filename pattern to match",
			},
			&cli.StringFlag{
				Name:    vendor,
				Aliases: []string{"v"},
				Value:   "github",
				Usage:   "a repository vendor name",
			},
			&cli.StringFlag{
				Name:    out,
				Aliases: []string{"o"},
				Value:   "json",
				Usage:   "an output format",
			},
		},
		Action: func(c *cli.Context) error {
			return RunFromContext(c)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func RunFromContext(c *cli.Context) error {

	var err error
	scanner := scanners.DockerFileScanner(c.String(input))
	if c.String(format) == "txt" {
		err = scanner.Read(readers.RemoteTxtReader(&http.Client{}))
		if err != nil {
			return err
		}
	}
	if c.String(vendor) == "github" {
		err = scanner.Validate(validators.GitHubValidator())
		if err != nil {
			return err
		}
		err = scanner.Query(github.Api(github.GoogleGitHubClient()))
		if err != nil {
			return err
		}
	}
	if c.String(pattern) == "Dockerfile" {
		err = scanner.Parse(parsers.DockerFileParser())
		if err != nil {
			return err
		}
	}
	if c.String(out) == "json" {
		err = scanner.Serialize(serializers.JsonSerializers())
		if err != nil {
			return err
		}
	}

	fmt.Println(scanner.GetData().Output)

	return err
}
