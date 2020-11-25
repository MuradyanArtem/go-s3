# Build
FROM golang:1.15-buster AS build

WORKDIR /app
ADD . .

ENV CGO_ENABLED=0

RUN go mod tidy
RUN go mod vendor
RUN go build -o bin/api ./src/cmd

# Enviroment
FROM alpine:latest

WORKDIR /app
COPY --from=build /app/bin/api .

ENTRYPOINT ["/app/api"]
EXPOSE 8080
