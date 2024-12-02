PHONY: migration migrate

migration:
	@read -p "Migration name : " name; \
	go run scripts/main.go migration $$name;

migrate:
	go run scripts/main.go migrate