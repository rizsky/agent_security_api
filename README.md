# BigAgent/SecurityAPI

BIG Agent's security-focused microservice.

## Service information

In `main.go`, there are `serviceName` and `serviceVersion` which represent this service general information. Should be updated whenever the version changes.

## Working on the service

We use [go-micro](https://github.com/micro/go-plugins) for handling our general microservice-related process and here's our tech stack for this service:

- Data storage: MongoDB, Redis
- Broker: Kafka
- Delivery: gRPC

### Environment variables

Listed on `.env.example`. Be cautious with `MICRO_` prefixed variables as it's related to go-micro configuration.