composeup:
	docker compose up -d

composedown:
	docker compose down

postgres:
	docker exec -it postgres psql -U postgres -b go_gin_jwt_oauth

run:
	go run cmd/main.go

.PHONY: composeup composedown run
