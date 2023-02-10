## Build
FROM golang:alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./ ./
COPY ./ ./

RUN go build -o /komiac-app ./cmd/main

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /komiac-app /komiac-app

COPY migrations ./migrations

ENTRYPOINT ["/komiac-app"]
