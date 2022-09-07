# syntax=docker/dockerfile:1
# Builder image
FROM golang:1.19-alpine AS build
WORKDIR /app

RUN apk add gcc musl-dev

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN go install golang.org/x/lint/golint@latest

COPY main.go ./

Run golint -set_exit_status

RUN go test

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Image
FROM alpine:3
WORKDIR /app

COPY assets ./assets
COPY templates ./templates
COPY --from=build /app/app ./

EXPOSE 8080

CMD ["./app"]
