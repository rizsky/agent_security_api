package main

import (
	"os"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/ralali/agent_security_api/helpers"
	"github.com/ralali/agent_security_api/infra/mongodb"
)

var mongoClient mongodb.Client
var mongoDatabase mongodb.Database

func initDatabase() {
	helpers.LogInfo("Initializing MongoDB connection")

	config := mongodb.Config{
		Host:         os.Getenv("MONGODB_HOST"),
		Port:         os.Getenv("MONGODB_PORT"),
		Username:     os.Getenv("MONGODB_USERNAME"),
		Password:     os.Getenv("MONGODB_PASSWORD"),
		DatabaseName: os.Getenv("MONGODB_DATABASE"),
	}

	client, err := mongodb.NewClient(config)
	if err != nil {
		helpers.LogError(err)
		panic(err)
	}

	mongoClient = client
	mongoDatabase = mongodb.NewDatabase(config, client)

	helpers.LogInfo("MongoDB connection initialized")
}

func initRepositories() {
	helpers.LogInfo("Initializing repositories")

	// TODO: init repositories here

	helpers.LogInfo("Repositories initialized")
}

func initServices() {
	helpers.LogInfo("Initializing services")

	// TODO: init services here

	helpers.LogInfo("Services initialized")
}

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
