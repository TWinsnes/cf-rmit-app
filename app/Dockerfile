# syntax=docker/dockerfile:1
# Builder image
FROM golang:1.16-alpine AS build
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Image
FROM alpine:3
WORKDIR /app

COPY assets ./assets
COPY templates ./templates
COPY --from=build /app/app ./

EXPOSE 8080

CMD ["./app"]
