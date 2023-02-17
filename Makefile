LINUX_AMD64 = CGO_ENABLED=0 GOOS=linux GOARCH=amd64

PROJECT_NAME = $(shell pwd | rev | cut -f1 -d'/' - | rev)

NEW_VERSION = $(shell expr $(CURRENT_VERSION) + 1 )

linter:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$GOPATH/bin v1.33.0

lint:
	golangci-lint run ./...

test:
	go test -covermode=count -coverprofile=count.out ./...

deps:
	go mod download
	go mod tidy

mockery:
	go install github.com/vektra/mockery/v2@latest

mock:
	mockery
	@rm -rf mocks
	mockery --dir api/service --all
	mockery --dir api/repository --all


build: 
	go build -o api/api ./api/main.go

build-linux:
	$(LINUX_AMD64) go build -o api/api ./api/main.go
	
build-image:
	@make deps
	@make build
	docker build -t $(PROJECT_NAME) .


# Running on windows set env to linux builder $Env:GOOS = "linux"; $Env:GOARCH = "amd64"; $Env:CGO_ENABLED = 0
local-start:
	@make build
	@docker-compose up

golang-migrate:
	go install github.com/golang-migrate/migrate/v4/cmd/migrate github.com/lib/pq github.com/hashicorp/go-multierror
	@go build -tags 'postgres' -o ${GOPATH}/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate


# Connection string parameters documentation: https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters
# Usage: DATABASE_URL=postgres://postgres:postgres@localhost:15432/logic-api?sslmode=disable make migrate
migrate: golang-migrate
	migrate -path migrations/ -database ${DATABASE_URL} up

# Usage: make migration name=my_migration
migration: golang-migrate
	migrate create -dir migrations/ -ext sql ${name}
