package main

import (
	"fmt"

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
		micro.Action(func(context *cli.Context) error {
			if context.Bool("serve") {
				if err := service.Run(); err != nil {
					fmt.Printf("[error] %s\n", err.Error())
					return err
				}
				return nil
			}

			if context.Bool("migrate") {
				fmt.Println("Running DB migrations...")
				return nil
			}

			if context.Bool("seed") {
				fmt.Println("Running data seeders...")
				return nil
			}

			return nil
		}),
	)
}
