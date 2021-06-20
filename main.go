package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "Scanner",
		Usage: "scans repositories for Dockerfiles and retrieves image information",
		Action: func(c *cli.Context) error {
			fmt.Println("Result")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
