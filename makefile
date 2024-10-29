PROTO_SRC_DIR := ./proto
PROTO_OUT_DIR := ./pb
GO_PACKAGE := github.com/napakornsk/cv-backend

PROTOC := protoc
PROTOC_GEN_GO := protoc-gen-go
PROTOC_GEN_GO_GRPC := protoc-gen-go-grpc

PROTO_FILES := $(wildcard $(PROTO_SRC_DIR)/*.proto)

all: proto

proto:
	$(PROTOC) --proto_path=$(PROTO_SRC_DIR) \
		--go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(PROTO_OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)


clean:
	rm -rf $(PROTO_OUT_DIR)/*.go

run:
	go run cmd/server/main.go

migrate:
	go run migration/migration.go

.PHONY: proto clean run migrate