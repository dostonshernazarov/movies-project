-include .env
export

.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	go build main.go && ./main

.PHONY: swag-gen
swag-gen:
	swag init -g core/app.go -o docs