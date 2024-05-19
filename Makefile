include .env

run:
	go run cmd/web/main.go

build:
	GOOS=linux GOARCH=amd64 go build -o fremont cmd/web/main.go

build-image:
	docker build -f ./deploy/Dockerfile -t malikilamalik/fremont . --no-cache

run-image:
	docker run -e POSTGRE_DB_NAME=$(POSTGRE_DB_NAME)  -e POSTGRE_DB_PORT=$(POSTGRE_DB_PORT) -e POSTGRE_DB_HOST=$(POSTGRE_DB_HOST) \
	-e POSTGRE_DB_USERNAME=$(POSTGRE_DB_USERNAME) \
	-e POSTGRE_DB_PASSWORD=$(POSTGRE_DB_PASSWORD) \
	-e POSTGRE_DB_PARAMS=$(POSTGRE_DB_PARAMS) \
	--network fremont_default \
	-p 8080:8080 malikilamalik/fremont:latest

migrate-dev:
	migrate -database "postgres://$(POSTGRE_DB_USERNAME):$(POSTGRE_DB_PASSWORD)@$(POSTGRE_DB_HOST):$(POSTGRE_DB_PORT)/$(POSTGRE_DB_NAME)?sslmode=disable" -path migrations -verbose up

rollback-dev:
	migrate -database "postgres://$(POSTGRE_DB_USERNAME):$(POSTGRE_DB_PASSWORD)@$(POSTGRE_DB_HOST):$(POSTGRE_DB_PORT)/$(POSTGRE_DB_NAME)?sslmode=disable" -path migrations -verbose down

drop-dev:
	migrate -database "postgres://$(POSTGRE_DB_USERNAME):$(POSTGRE_DB_PASSWORD)@$(POSTGRE_DB_HOST):$(POSTGRE_DB_PORT)/$(POSTGRE_DB_NAME)?sslmode=disable" -path migrations -verbose drop