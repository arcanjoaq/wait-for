#!/usr/bin/make -f
clean:
	rm app.linux-amd64
lint:
	find . -type f -name '*.go' -not -path "vendor*" -exec gofmt -w {} \;
build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o wait-for
run:
	go run main.go
docker:
	docker build -t arcanjoqueiroz/golang-sandbox .
