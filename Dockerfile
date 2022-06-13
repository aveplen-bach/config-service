FROM golang:1.18

RUN mkdir -p /go/src/github.com/aveplen-bach/config_service

WORKDIR /go/src/github.com/aveplen-bach/config_service

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . ./

RUN go build -o bin/auth cmd/main.go

ENTRYPOINT [ "go", "run", "cmd/main.go" ]