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
	app.Usage = "GRPC service client"
	app.EnableBashCompletion = true
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "server-uri",
			Aliases:  []string{"s"},
			EnvVars:  []string{"SERVER_URI"},
			Required: true,
			Usage:    "grpc server connection",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:   "fetch",
			Usage:  "Fetch csv records",
			Action: fetch,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "fetch-uri",
					Aliases:  []string{"f"},
					EnvVars:  []string{"FETCH_URI"},
					Required: true,
					Usage:    "fetch uri",
				},
				&cli.StringFlag{
					Name:     "separator",
					Aliases:  []string{"s"},
					EnvVars:  []string{"SEPARATOR"},
					Value:    ";",
					Required: false,
					Usage:    "csv records separator",
				},
				&cli.BoolFlag{
					Name:     "skip-header",
					Aliases:  []string{"sh"},
					EnvVars:  []string{"SKIP_HEADER"},
					Required: false,
					Usage:    "skip csv header",
				},
			},
		},
		{
			Name:   "show",
			Usage:  "List records",
			Action: list,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     "sorting",
					Aliases:  []string{"s"},
					Required: false,
					Value:    "name",
					Usage:    "sorting field",
				},
				&cli.StringFlag{
					Name:     "order",
					Aliases:  []string{"o"},
					Required: false,
					Value:    "asc",
					Usage:    "sorting order",
				},
			},
		},
	}
	return app
}
