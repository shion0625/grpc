.PHONY: test coverage

test:
	ginkgo -v ./...

# テストカバレッジの計測
coverage:
	ginkgo -cover --race --output-dir=test/coverage -coverprofile=coverage.out ./...
	go tool cover -html=test/coverage/coverage.out -o test/coverage/coverage.html

# コードの整形
format:
	go fmt ./...
	goimports -w .

# 静的解析ツールの実行
lint: format
	golangci-lint run ./...

# 依存関係のインストール
deps:
	go mod tidy

# すべてのテストと解析を実行
all: deps lint test coverage

sqlboiler-commandservice:
	sqlboiler mysql -c internal/app/commandservice/infra/sqlboiler/config/database.toml -o internal/app/commandservice/infra/sqlboiler/models -p models --no-tests --wipe

db-key: 
	mkdir -p ddl/ssl
	openssl req -newkey rsa:2048 -nodes -keyout ddl/ssl/server-key.pem -x509 -out ddl/ssl/server-cert.pem
	openssl req -newkey rsa:2048 -nodes -keyout ddl/ssl/ca-key.pem -x509 -out ddl/ssl/ca-cert.pem