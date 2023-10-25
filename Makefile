.PHONY: build
build:
	go build -v ./app/cmd

.PHONY: test
test:
	go test -v -timeout 30s ./...

migrations_up:
	migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

.DEFAULT_GOAL := build

run:
	docker run -d -p 80:3000 --rm --name rest-api rest-api:env
stop:
	docker stop rest-api