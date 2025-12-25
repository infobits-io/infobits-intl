.PHONY: all generate generate-dart generate-typescript generate-go validate extract test clean lint lint-generator lint-go deploy deploy-typescript deploy-dart deploy-typescript-dry deploy-dart-dry

# Default target
all: generate

# Generate all targets
generate:
	cd generator && go run . generate --all --config ../generator.yaml --data ../data --templates ./templates

# Generate specific targets
generate-dart:
	cd generator && go run . generate --target dart --config ../generator.yaml --data ../data --templates ./templates

generate-typescript:
	cd generator && go run . generate --target typescript --config ../generator.yaml --data ../data --templates ./templates

generate-go:
	cd generator && go run . generate --target go --config ../generator.yaml --data ../data --templates ./templates

# Validate data files
validate:
	cd generator && go run . validate --data ../data

# Extract data from existing Dart files (migration helper)
extract:
	cd generator && go run . extract --source ../lib/src --output ../data/core

# Run tests for all packages
test: test-dart test-typescript test-go

test-dart:
	cd packages/dart && flutter test

test-typescript:
	cd packages/typescript && npm test

test-go:
	cd packages/go && go test ./...

# Clean generated files
clean:
	rm -f generator/intlgen
	rm -rf packages/dart/lib/src/*.g.dart
	rm -rf packages/dart/lib/src/i18n/
	rm -rf packages/typescript/src/*.ts
	rm -rf packages/typescript/src/i18n/
	rm -rf packages/go/*.go
	rm -rf packages/go/i18n/

# Lint all Go code
lint: lint-generator lint-go

lint-generator:
	cd generator && golangci-lint run ./...

lint-go:
	cd packages/go && golangci-lint run ./...

# Deploy packages
deploy: deploy-typescript deploy-dart

deploy-typescript:
	cd packages/typescript && npm run build && npm publish --access public

deploy-dart:
	cd packages/dart && dart pub publish --force

# Dry run (test without publishing)
deploy-typescript-dry:
	cd packages/typescript && npm run build && npm publish --access public --dry-run

deploy-dart-dry:
	cd packages/dart && dart pub publish --dry-run