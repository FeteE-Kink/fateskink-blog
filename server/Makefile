MIGRATIONS_FOLDER = $(PWD)/app/database/migrations
DB_USERNAME=root
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=fateskink
DB_PASSWORD=
DB_URL=mysql://$(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_DATABASE)?parseTime=true&loc=Local

POSTGRES_DATABASE_NAME ?= fateskink
POSTGRES_HOST ?= localhost
POSTGRES_PORT ?= 5432
POSTGRES_USERNAME ?= admin
POSTGRES_PASSWORD ?= 12345678
POSTGRES_SSL_MODE ?= disable

DB_URL=postgres://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DATABASE_NAME)?sslmode=$(POSTGRES_SSL_MODE)

cmd-exists-%:
	@hash $(*) > /dev/null 2>&1 || \
	   (echo "ERROR: '$(*)' Must be installed  and available on your PATH."; exit 1)

db.create: cmd-exists-migrate
	createdb --username=$(POSTGRES_USERNAME) --owner=$(POSTGRES_USERNAME) ${POSTGRES_DATABASE_NAME}

db.drop: cmd-exists-migrate
	dropdb --username=$(POSTGRES_USERNAME) ${POSTGRES_DATABASE_NAME}

db.migrate.rollback: cmd-exists-migrate
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DB_URL)" --verbose down ${STEP}

db.migrate: cmd-exists-migrate
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DB_URL)" --verbose up

db.migrate.force: cmd-exists-migrate
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DB_URL)" --verbose force $(VERSION)

db.seed: cmd-exists-migrate
	go run app/pkg/initializers/seeds/tasks/main.go
