FROM golang:1.16 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct"

WORKDIR /app-upgrade

COPY . .

RUN go mod download

RUN go build -o main main.go

EXPOSE 9090

ENTRYPOINT ["./main"]
