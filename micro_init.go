package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
)

func createService() micro.Service {
	return micro.NewService(
		micro.Name(serviceName),
		micro.Version(serviceVersion),
		micro.Flags(
			&cli.BoolFlag{
				Name:  "migrate",
				Usage: "Run DB migrations",
			},
			&cli.BoolFlag{
				Name:  "seed",
				Usage: "Run data seeders",
			},
		),
	)
}
