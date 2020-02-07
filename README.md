# BIGAgent/SecurityAPI

BIG Agent's security-focused microservice.

## Service information

In `main.go`, there are `serviceName` and `serviceVersion` which represent this service general information. Should be updated whenever the version changes.

## Working on the service

We use Go 1.13.0 with go modules.

We use [go-micro](https://github.com/micro/go-plugins) for handling our general microservice-related process and here's our tech stack for this service:

- Data storage: MongoDB, Redis
- Broker: Kafka
- Delivery: gRPC

### Setup local environment

We provide `docker-compose.yml` which will launch tech stack needed to develop this service locally:

1. ZooKeper
2. Kafka
3. Kafdrop (Kafka UI) (accessible via `http://localhost:9000`)
4. Consul (accessible via `http://localhost:8500/ui`)

But for Kafka, we need to add 1 record to our `/etc/hosts` (Unix):

```
# Ralali
127.0.0.1 kafka
```

### Environment variables

Listed on `.env.example`. Be cautious with `MICRO_` prefixed variables as it's related to go-micro configuration.