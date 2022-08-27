FROM golang:1.19.0-alpine3.16

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./
RUN : && \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64 \
  go build -a -o app . && \
  :

EXPOSE 8080