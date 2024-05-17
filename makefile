up:
	@echo "Starting docker images..."
	docker-compose up -d
	@echo "Docker started!"

up_build:
	@echo Building docker containers...
	docker-compose up -d --build
	@echo Done!

down:
	@echo Stopping docker images...
	docker-compose down
	@echo Done!

swag:
	@echo Updating swagger documentation...
	swag init -g cmd/api/main.go
	@echo Done!

restart: down up_build

gotest:
	go test ./...


lint:
	golangci-lint run

lgbt: lint gotest