DOCKER ?= docker
GO ?= go

# Tooling expects a fully qualified path to the plugins directory.
BUF_PLUGINS_DIR ?= $(shell pwd)/.buf-plugins

DIR := ${CURDIR}
OUT_DIR ?= .out
DIST_DIR = .dist
BIN_DIR = ${OUT_DIR}/bin
CERTS_DIR = certs

# LDFLAGS += $(shell ./tools/build/go-ldflags.sh)
# GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o "$(OUT_DIR)/iam-darwin-amd64/iam" .
# GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o "$(OUT_DIR)/ias-darwin-arm64/iam" .

.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 go build -o "$(OUT_DIR)/iam-darwin-amd64/iam" .
	GOOS=darwin GOARCH=arm64 go build -o "$(OUT_DIR)/iam-darwin-arm64/iam" .

.PHONY: clean
clean:
	rm -rf $(OUT_DIR)
	rm -rf $(CERTS_DIR)
	mkdir -p $(CERTS_DIR)

.PHONY: protoc
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/grpc/v1/buzz/*.proto

.PHONY: certs
certs:
	openssl req -new -newkey rsa:4096 -x509 -sha256 -days 365 -nodes -out certs/server.crt -keyout certs/server.key -subj "/C=US/ST=WA/L=Seattle/O=IAM/OU=IAM/CN=dissentr.io" -addext "subjectAltName=DNS:dissentr.io,DNS:localhost"