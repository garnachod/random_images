
build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/random ./cmd/random

docker-build:
	@echo "Using docker for $*"
	@docker-compose -f docker-compose-build.yml build
	@docker-compose -f docker-compose-build.yml run --rm app make build
	@docker-compose -f docker-compose-build.yml down

docker-run:
	@echo "Using docker for $*"
	@docker-compose build
	@docker-compose down
	@docker-compose up