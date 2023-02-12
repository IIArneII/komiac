# Komiac

Microservice for work with applications for subsidized medicines.

## Requirements

- go 1.19

## Configurations

Create a `.env` file and fill it with the environment variables you want to use, following the pattern of the `.env.sample` file. Also you can create necessary environment variables without creating a file `.env`.

**Important:** environment variables in `.env` file are used for docker configuration.

## Build and run the application

Loading the required libraries:
```bash
go mod download
```

Build and run the application:
```bash
go build -o ./main ./cmd/main
./main
```

## Docker

```bash
docker-compose build
docker-compose up -d
```
