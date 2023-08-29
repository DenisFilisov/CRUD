build:
	go mod download && go build ./cmd/main.go

run: build
	docker-compose up --remove-orphans

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:postgres@0.0.0.0:5432/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go
