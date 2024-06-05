MIGRATIONS_FOLDER = $(PWD)/database/migrations
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

db.create:
	createdb --username=$(POSTGRES_USERNAME) --owner=$(POSTGRES_USERNAME) ${POSTGRES_DATABASE_NAME}

db.drop:
	dropdb --username=$(POSTGRES_USERNAME) ${POSTGRES_DATABASE_NAME}

db.migrate.rollback:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DB_URL)" --verbose down ${STEP}

db.migrate:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DB_URL)" --verbose up

db.migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DB_URL)" --verbose force $(VERSION)
