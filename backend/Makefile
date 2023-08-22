dev:
	@docker-compose up --build

stop:
	@docker stop $$(docker ps -q)

run:
	@go install github.com/cespare/reflex@latest
	@reflex -r '\.go$$' -s go run main.go

db_exec:
	@docker-compose exec mysql /bin/bash
	# mysql -u user -ppassword -h mysql -P 3306

test:
	@go test -v ./...

lint:
	@golangci-lint run

fix:
	@golangci-lint run --fix