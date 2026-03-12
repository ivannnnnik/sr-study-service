DB_DSN ?= postgres://postgres:postgres@localhost:5432/sr_study_db?sslmode=disable

.PHONY: migrate-up migrate-down migrate-create

migrate-up:
	migrate -path migrations -database "$(DB_DSN)" up

migrate-down:
	migrate -path migrations -database "$(DB_DSN)" down

migrate-create:
	@read -p "Migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name
