build:
	@go build -o bin/api

run: build
	@./bin/api

seed: 
	@go run seeders/db_seeder.go

test: 
	@go test -v ./tests/