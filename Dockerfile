FROM golang:1.18 as builder

RUN mkdir -p /go/src/github.com/aveplen-bach/config-service

WORKDIR /go/src/github.com/aveplen-bach/config-service

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o /bin/config-service \
    /go/src/github.com/aveplen-bach/config-service/cmd/main.go

FROM alpine:3.15.4 as runtime

COPY --from=builder /bin/config-service /bin/config-service

ENTRYPOINT [ "/bin/config-service" ]%