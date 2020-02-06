package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
)

const serviceName = "BIGAgent/SecurityAPI"
const serviceVersion = "1.0.0"

func main() {
	godotenv.Load()

	service := createService()

	service.Init(
		micro.Action(func(ctx *cli.Context) error {
			if ctx.Bool("migrate") {
				// TODO: implement migration here (maybe)
				fmt.Println("Running DB migrations...")
				os.Exit(1)
			}

			if ctx.Bool("seed") {
				// TODO: implement seeding here (maybe)
				fmt.Println("Running data seeders...")
				os.Exit(1)
			}

			return nil
		}),
	)

	if err := service.Run(); err != nil {
		fmt.Println("[error]", err.Error())
	}
}
