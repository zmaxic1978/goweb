.SILENT:

build:
	docker-compose up -d --build

run:
	docker pull postgres &&
    docker run --name=todo-db -e POSTGRES_PASSWORD='12345' -p 5436:5432 -d --rm postgres &&
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5436/postgres?sslmode=disable' up &&
	go run cmd/main.go
