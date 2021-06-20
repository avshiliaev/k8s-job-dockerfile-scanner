package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
	"redhat-sre-task-dockerfile-scanner/src/api/github"
	"redhat-sre-task-dockerfile-scanner/src/parsers"
	"redhat-sre-task-dockerfile-scanner/src/readers"
	scanners "redhat-sre-task-dockerfile-scanner/src/scanner"
	"redhat-sre-task-dockerfile-scanner/src/validators"
	"redhat-sre-task-dockerfile-scanner/src/writers"
)

func main() {

	// ./scanner -i http://github.com/ -p Dockerfile -v github -o json
	app := &cli.App{
		Name:  "Scanner",
		Usage: "scans repositories for Dockerfiles and retrieves image information",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "input",
				Aliases:  []string{"i"},
				Value:    "",
				Usage:    "link to a txt input file",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "txt",
				Usage:   "input format",
			},
			&cli.StringFlag{
				Name:    "pattern",
				Aliases: []string{"p"},
				Value:   "Dockerfile",
				Usage:   "a filename pattern to match",
			},
			&cli.StringFlag{
				Name:    "vendor",
				Aliases: []string{"v"},
				Value:   "github",
				Usage:   "a repository vendor name",
			},
			&cli.StringFlag{
				Name:    "out",
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
	inputLink := c.String("input")
	inputFormat := c.String("format")
	repoVendor := c.String("vendor")
	fileNamePattern := c.String("pattern")
	outputFormat := c.String("out")

	scanner := scanners.DockerFileScanner(inputLink)
	if inputFormat == "txt" {
		scanner.Read(readers.RemoteTxtReader(&http.Client{}))
	}
	if repoVendor == "github" {
		scanner.Validate(validators.GitHubValidator())
		scanner.Query(github.Api(github.GoogleGitHubClient()))
	}
	if fileNamePattern == "Dockerfile" {
		scanner.Parse(parsers.DockerFileParser())
	}
	if outputFormat == "json" {
		scanner.Write(writers.JsonStdWriter())
	}

	return nil
}
