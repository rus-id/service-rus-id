.PHONY:build
build:
	echo "building.."
	go build -o bin/service-rus-id ./cmd/service/main.go

.PHONY:run
run:
	echo "running.."
	go run -race ./cmd/service/main.go

.PHONY:test
test:
	echo "testing.."
	go test -v -cover ./...

.DEFAULT_GOAL := run