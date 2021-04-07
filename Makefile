#!/usr/bin/make -f
clean:
	rm wait-for
lint:
	find . -type f -name '*.go' -not -path "vendor*" -exec gofmt -w {} \;
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o wait-for
