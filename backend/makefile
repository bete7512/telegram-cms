include .env
install_deps:
	go get -u github.com/swaggo/swag/cmd/swag\
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
prepare_docs:
	swag init -g ./cmd/main.go -o ./docs
create_migration:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations $$name
migrate_up_all:
	migrate -database $(DB_URI) -path ./migrations up
migrate_down_all:
	migrate -path ./migrations -database $(DB_URI) down
migrate_up_last:
	migrate -database $(DB_URI) -path ./migrations up 1
migrate_down_last:
	migrate -path ./migrations -database $(DB_URI) down 1
migrate_up_to:
	@read -p "Enter migration version: " version; \
	migrate -path ./migrations -database $(DB_URI) goto $$version
migrate_down_to:
	@read -p "Enter migration version: " version; \
	migrate -path ./migrations -database $(DB_URI) goto $$version
