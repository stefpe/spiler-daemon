install:
	go get -v -d -t ./...

run:
	go run main.go -address=127.0.0.1:9001 -endpoint=http://127.0.0.1:8000/server.php

run-server:
	php -S localhost:8000

build:
	go build
