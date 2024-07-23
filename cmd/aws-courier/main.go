package main

import (
	"fmt"
	"os"

	"github.com/gemcdaniel/aws-courier/aws"
	"github.com/gemcdaniel/aws-courier/http"
	cli "github.com/jawher/mow.cli"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	app := cli.App("awsprofile", "A restful API to serve aws credentials")
	app.Version("v version", "v1.0.0")

	// Define Global Flags
	var (
		port = app.StringOpt("p port", "8080", "The port to run the API")
		file = app.StringOpt("f file", "", "A non-defaulted location of the credentials file to use")
	)

	app.Action = func() {
		credentialsService := aws.NewCredentialsService(*file)

		router := http.NewHandler()
		router.CredentialsService(credentialsService)

		svr := http.NewServer()
		svr.UseHandler(router)

		fmt.Println("Running server")
		svr.Run(":" + *port)
	}

	return app.Run(os.Args)
}
