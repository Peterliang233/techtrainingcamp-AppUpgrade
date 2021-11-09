echo "begin project..."

export GO111MODULE=on

export GOPROXY=https://goproxy.cn,direct

go mod tidy

go run main.go