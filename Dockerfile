FROM golang:1.22.5

WORKDIR /home/finux/github.com/drivehino-trivia-restapi

RUN go install github.com/air-verse/air@latest

COPY . .
RUN go mod tidy
