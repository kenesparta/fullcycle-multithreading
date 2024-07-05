run:
	go run main.go

.PHONY: init
init:
	go mod tidy

.PHONY: build
build:
	GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o server ./cmd/api
