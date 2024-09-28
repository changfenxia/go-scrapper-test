db_up:
    docker-compose --env-file .env up -d

db_down:
    docker-compose --env-file .env down

db_migration:
    @if [ -z "$(name)" ]; then echo "Использование: make db_migration name=name_of_migration_file"; exit 1; fi
    go run github.com/pressly/goose/v3/cmd/goose -dir ./migrations create $(name) sql