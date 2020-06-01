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
	go clean -testcache
	go test -v -cover -tags=entities ./...

.PHONY:proto
proto:
	protoc --go_out=plugins=grpc,paths=import:. api/service/*.proto

.DEFAULT_GOAL := run