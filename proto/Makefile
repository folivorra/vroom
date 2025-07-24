BUF_VERSION                     = v1.54.0
PROTOC_GEN_GO_VERSION           = v1.36.6
PROTOC_GEN_GO_GRPC_VERSION      = v1.5.1
PROTOC_GEN_GRPC_GATEWAY_VERSION = v2.26.3
PROTOC_GEN_VALIDATE_VERSION     = v1.2.1

ROOT_DIR 					= $(patsubst %/,%,$(dir $(realpath $(lastword $(MAKEFILE_LIST)))))
BIN_DIR          			= $(ROOT_DIR)/bin
BUF              			= $(BIN_DIR)/buf
GEN_DIR						= /gen
PROTOC_GEN_GO    			= $(BIN_DIR)/protoc-gen-go
PROTOC_GEN_GO_GRPC 			= $(BIN_DIR)/protoc-gen-go-grpc
PROTOC_GEN_GRPC_GATEWAY     = $(BIN)/protoc-gen-grpc-gateway
PROTOC_GEN_VALIDATE         = $(BIN)/protoc-gen-validate

.PHONY: all gen clean

all: $(BUF) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC) $(PROTOC_GEN_GRPC_GATEWAY) $(PROTOC_GEN_VALIDATE)

# ---------- buf --------------------------------------------------------------
$(BUF):
	@mkdir -p $(BIN_DIR)
	@echo "⏬  downloading buf $(BUF_VERSION)"
	@GOBIN=$(BIN_DIR) go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)
	@echo "✓ buf installed → $(BUF)"

# ---------- protoc-плагины ----------------------------------------------------
$(PROTOC_GEN_GO):
	@mkdir -p $(BIN_DIR)
	@echo "⏬  installing protoc-gen-go $(PROTOC_GEN_GO_VERSION)"
	@GOBIN=$(BIN_DIR) go install \
	    google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)

$(PROTOC_GEN_GO_GRPC):
	@mkdir -p $(BIN_DIR)
	@echo "⏬  installing protoc-gen-go-grpc $(PROTOC_GEN_GO_GRPC_VERSION)"
	@GOBIN=$(BIN_DIR) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

$(PROTOC_GEN_GRPC_GATEWAY):
	@mkdir -p $(BIN_DIR)
	@echo "⏬  installing protoc-gen-grpc-gateway $(PROTOC_GEN_GRPC_GATEWAY_VERSION)"
	@GOBIN=$(BIN_DIR) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@$(PROTOC_GEN_GRPC_GATEWAY_VERSION)

# ---------- генерация через buf ----------------------------------------------
gen-buf: $(BUF) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC)
	@echo "⚙️  buf generate"
	@$(BUF) generate
	@echo "✓ protobuf code generated"

# ---------- генерация через protoc ----------------------------------------------
gen-protoc:
	@mkdir -p $(GEN_DIR)
	@docker run --rm -v $(ROOT_DIR):/defs namely/protoc-all -f dumpd/v1/dumpd.proto -l go -o /defs/gen

# ---------- очистка -----------------------------------------------------------
clean:
	@rm -rf $(BIN_DIR)
	@echo "🧹  bin/ очищены"