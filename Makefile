dev:
	@docker-compose up --build

stop:
	@docker stop $$(docker ps -q)

run:
	@go install github.com/cespare/reflex@latest
	@reflex -r '\.go$$' -s go run main.go

test:
	@go test -v ./...