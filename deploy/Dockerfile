# syntax=docker/dockerfile:1

# build
FROM golang:1.22.0 as build
WORKDIR /usr/src/app
ARG CGO_ENABLED=0
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o fremont cmd/web/main.go

# deploy
FROM alpine:3.19

WORKDIR /app
COPY --from=build usr/src/app/fremont .

EXPOSE 8080
ENTRYPOINT [ "./fremont" ]