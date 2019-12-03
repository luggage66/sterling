GOOS=linux
GOARCH=amd64

build:
	go mod vendor
	go build -mod vendor -buildmode=exe -o bin/foo-thing main.go