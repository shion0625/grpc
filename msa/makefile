.PHONY: test coverage

test:
	ginkgo -v ./...

# テストカバレッジの計測
coverage:
	ginkgo -cover --race --output-dir=test/coverage -coverprofile=coverage.out ./...
	go tool cover -html=test/coverage/coverage.out -o test/coverage/coverage.html

# 静的解析ツールの実行
lint:
	golangci-lint run ./...

# 依存関係のインストール
deps:
	go mod tidy

# すべてのテストと解析を実行
all: deps lint test coverage