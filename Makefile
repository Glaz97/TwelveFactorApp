.PHONY: swag start

swag:
	swag fmt
	swag init -g internal/handler/handler.go --output pkg/docs

mongo:
	docker compose -f test/docker-compose.yaml -p twelvefactorapp up -d

mongo-down:
	docker compose -f test/docker-compose.yaml -p twelvefactorapp down

test: mongo
	go test -v -race ./...

start: mongo
	go run cmd/twelvefactorapp/main.go config  > configs/twelvefactorapp.yaml