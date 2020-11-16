package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/trezorg/atlant/pkg/server"
	"github.com/urfave/cli/v2"
)

func start(c *cli.Context) error {
	logLevel := c.String("log-level")
	host := c.String("address")
	port := c.Int("port")
	mongoURI := c.String("mongo-uri")

	ctx, done := context.WithCancel(context.Background())
	defer done()

	return server.Start(ctx, mongoURI, host, port, logLevel)
}

func main() {
	app := prepareArgs()
	if err := app.Run(os.Args); err != nil {
		logrus.Errorf("Cannot initialize application: %+v", err)
		os.Exit(1)
	}
}
