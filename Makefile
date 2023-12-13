MAKEFLAGS += --silent
.DEFAULT_GOAL := dev

dev:
	watchexec -r -e go go run main.go

.PHONY: dev

auto-migrate:
	go run scripts/migrations/migrations.go

.PHONY: auto-migrate

auto-seed:
	go run scripts/seeds/seeds.go

.PHONY: auto-seed

db-down:
	docker-compose down && docker volume rm template_postgres_volume 2> /dev/null

.PHONY: db-down

db-up:
	docker-compose up -d

.PHONY: db-up

reset-db:
	make db-down; make db-up && sleep 2 && make auto-migrate && make auto-seed

.PHONY: reset-db