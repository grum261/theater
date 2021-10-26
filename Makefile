include .env.dev

migrate-up:
	migrate -path internal/migrations -database=pgx://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable up

run:
	go run cmd/rest_server/main.go -env .env.dev