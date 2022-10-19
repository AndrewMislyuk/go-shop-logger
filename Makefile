BINARY=grpcApp

## generate_audit: generate protobuf audit
generate_audit:
	@echo "Generation protobuf..."
	protoc -I proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. proto/audit.proto
	@echo "Complete"

## build application
build:
	@echo "Building binary..."
	env GOOS=linux CGO_ENABLED=0 go build -o ${BINARY} ./cmd
	@echo "Done!"

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d --build
	@echo "Docker images started!"

## down: stop docker compose
down:
	@echo "Stop all containers docker compose..."
	docker-compose down
	@echo "Done!"

## build_up: compile and build app
build_up: build up
	@echo "Delete binary app"
	rm -f -r grpcApp
	@echo "Done!"