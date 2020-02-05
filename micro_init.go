package main

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func createService() micro.Service {
	registry := consul.NewRegistry()

	return micro.NewService(
		micro.Name(serviceName),
		micro.Version(serviceVersion),
		micro.Registry(registry),
		micro.Flags(
			&cli.BoolFlag{
				Name:  "serve",
				Usage: "Run the microservice",
			},
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
