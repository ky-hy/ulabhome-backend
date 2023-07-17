.DEFAULT_GOAL := help

.PHONY: bl
bl: ## Build docker image to local development
	docker build --no-cache --target deploy ./

.PHONY: bu
bu: ## Build docker image and up container
	docker compose up -d --build

.PHONY: serve
serve: ## serve with air 
	docker compose exec app air

.PHONY: in
in: ## Appのコンテナに入る
	docker compose exec app sh

.PHONY: up
up: ## Do docker compose up with hot reload
	docker compose up -d

.PHONY: down
down: ## Do docker compose down
	docker compose down

.PHONY: logs
logs: ## Tail docker compose logs
	docker compose logs -f app

.PHONY: ps
ps: ## Check container status
	docker compose ps

.PHONY: dry-migrate
dry-migrate: ## Try migration
	mysqldef -u ${DB_USER} -p ${DB_PASSWORD} -h ${DB_HOST} -P ${DB_PORT} ${DB_NAME} --dry-run < ./_tools/mysql/schema.sql

.PHONY: migrate
migrate:  ## Execute migration
	mysqldef -u ${DB_USER} -p ${DB_PASSWORD} -h ${DB_HOST} -P ${DB_PORT} ${DB_NAME} < ./_tools/mysql/schema.sql

.PHONY: seed
seed: ## seed data to db
	mysql ${DB_NAME} -h ${DB_HOST} -u ${DB_USER} -p${DB_PASSWORD} < ./_tools/mysql/seed.sql 

.PHONY: format
format: ## フォーマット
	gofmt -l -s -w .
	goimports -w -l .

.PHONY: lint
lint: ## リンター(golangci-lint)
	golangci-lint run

.PHONY: test
test: ## テスト
	go test -cover -race -shuffle=on ./...

.PHONY: mc
mc: ## make coverage カバレッジファイル作成（コンテナ側）
	go test -cover ./... -coverprofile=cover.out
	go tool cover -html=cover.out -o tmp/cover.html
	rm cover.out

.PHONY: wc
wc: ## watch coverage カバレッジを見る（ホスト側）
	open ./tmp/cover.html

.PHONY: swag
swag: ## swaggorファイル作成
	swag init --outputTypes yaml,go

.PHONY: help
help: ## Show options
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
