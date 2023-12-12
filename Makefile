MAKEFLAGS += --silent
.DEFAULT_GOAL := auto-migrate

auto-migrate:
	go run scripts/migrations/migrations.go

.PHONY: auto-migrate

auto-seed:
	go run scripts/seeds/seeds.go

.PHONY: auto-seed

db-down:
	docker-compose down && docker volume rm template_postgres_volume

.PHONY: db-down

db-up:
	docker-compose up -d

.PHONY: db-up

reset-db:
	make db-down && make db-up && sleep 5 && make auto-migrate

.PHONY: reset-db