# Phonebook

Project to pratice golang development, based on [Juju's Challenge - Agenda em Go](https://www.notion.so/Juju-s-Challenge-Agenda-em-Go-972dce60444f46f5bd681e4c9c1c614d).


# Pre requisites

- [go](https://golang.org/doc/install)
- [docker](https://docs.docker.com/engine/install/)
- [docker-compose](https://docs.docker.com/compose/install/)

## Running the project

To run the project you just need to run:

```
docker-compose up -d
```

This command will run an instance of MongoDB, Grafana, Prometheus and the application Phonebook.

## Docs

An postman collection was provided with the requests to using the application. You can import using the file `postman_collection.json`

## Metrics

The application exposes metrics for prometheus on `/metrics` endpoint.

You can visualize theese metrics using Grafana, after ran the project you can access `localhost:3000` and visualize the Phonebook dashboard.

The default credentials for grafana are:

```
username: admin
password: admin
```

## Testing

### Unit tests

To run the unit tests you need to run the following command:

```
go test ./internal/...
```

### End to end tests

To run the end to end tests, make sure that the containers are not running executing:

```
docker-compose down -v
```

Then, run the e2e script execute:

```
./e2e_test.sh
```
