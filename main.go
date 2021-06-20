package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	// ./scanner -input http://github.com/ -pattern Dockerfile
	app := &cli.App{
		Name:  "Scanner",
		Usage: "scans repositories for Dockerfiles and retrieves image information",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "input",
				Value: "",
				Usage: "link to a txt input file",
			},
			&cli.StringFlag{
				Name:  "pattern",
				Value: "Dockerfile",
				Usage: "a filename pattern to match",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println(c.String("input"))
			fmt.Println(c.String("pattern"))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
