
build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/random ./cmd/random

docker-build:
	@echo "Using docker for $*"
	@docker-compose -f docker-compose-build.yml build
	@docker-compose -f docker-compose-build.yml run --rm app make build
	@docker-compose -f docker-compose-build.yml down

build-mocks:
	mockgen -source=internal/user/provider.go -destination internal/user/mock_provider.go -package user
	mockgen -source=internal/image/provider.go -destination internal/image/mock_provider.go -package image

test:
	go test -cover ./internal/user ./internal/image

docker-test:
	@echo "Using docker for $*"
	@docker-compose -f docker-compose-build.yml build
	@docker-compose -f docker-compose-build.yml run --rm app make test
	@docker-compose -f docker-compose-build.yml down

docker-run:
	@echo "Using docker for $*"
	@docker-compose build
	@docker-compose down
	@docker-compose up