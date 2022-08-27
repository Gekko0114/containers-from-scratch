FROM golang:1.19.0-alpine3.16

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN go build -o /container-from-scratch

EXPOSE 8080