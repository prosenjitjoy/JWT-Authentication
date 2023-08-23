include .env

create_container:
	podman run --name ${CONTAINER_NAME} -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASS} -p ${DB_PORT}:5432 -d postgres

create_database:
	podman exec -it ${CONTAINER_NAME} createdb --username=${DB_USER} ${DB_NAME}

delete_database:
	podman exec -it ${CONTAINER_NAME} dropdb --username=${DB_USER} ${DB_NAME}

open_database:
	podman exec -it ${CONTAINER_NAME} psql -U ${DB_USER} -d ${DB_NAME}

create_migrations:
	migrate create -ext sql -dir database/migrations -seq create_tables

migrate_up:
	migrate -database ${MIGRATE_URL} -path database/migrations up

migrate_down:
	migrate -database ${MIGRATE_URL} -path database/migrations down

.PHONY: create_container create_database delete_database open_database migrate_up migrate_down