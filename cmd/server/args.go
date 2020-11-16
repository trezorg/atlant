package main

import (
	"github.com/trezorg/atlant/cmd"
	"github.com/urfave/cli/v2"
)

func prepareArgs() *cli.App {
	app := cli.NewApp()
	app.Version = cmd.Version
	app.HideHelp = false
	app.HideVersion = false
	app.Authors = []*cli.Author{{
		Name:  "Igor Nemilentsev",
		Email: "trezorg@gmail.com",
	}}
	app.Usage = "GRPC service"
	app.EnableBashCompletion = true
	app.Action = start
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "address",
			Aliases:  []string{"a"},
			EnvVars:  []string{"HOST"},
			Value:    "0.0.0.0",
			Required: false,
			Usage:    "Listening address",
		},
		&cli.StringFlag{
			Name:     "log-level",
			Aliases:  []string{"l"},
			EnvVars:  []string{"LOG_LEVEL"},
			Value:    "INFO",
			Required: false,
			Usage:    "Log level",
		},
		&cli.StringFlag{
			Name:     "mongo-uri",
			Aliases:  []string{"c"},
			EnvVars:  []string{"MONGO_URI"},
			Required: true,
			Usage:    "mongodb connection URI",
		},
		&cli.IntFlag{
			Name:     "port",
			Aliases:  []string{"p"},
			EnvVars:  []string{"PORT"},
			Value:    10000,
			Required: false,
			Usage:    "Listening port",
		},
	}
	return app
}
