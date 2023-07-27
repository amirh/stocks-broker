PROTOC_GO_VERSION ?= v1.28.1
PROTOC_GRPC_GO_VERSION ?= v1.2.0

BIN = $(CURDIR)/.bin

PROTOC = $(or ${PROTOC_BIN},protoc)
GO := go

get-protoc-deps:
	@GOBIN=$(BIN) $(GO) install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GO_VERSION)
	@GOBIN=$(BIN) $(GO) install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GRPC_GO_VERSION)

.PHONY: protocompile
protocompile: get-protoc-deps
	@ if ! which $(PROTOC) > /dev/null; then \
			echo "error: protoc not installed" >&2; \
			exit 1; \
		fi
	\
	cd proto && \
	  $(PROTOC) --plugin=protoc-gen-go=$(BIN)/protoc-gen-go --plugin=protoc-gen-go-grpc=$(BIN)/protoc-gen-go-grpc \
	    --proto_path=. --go_opt=paths=source_relative --go_out=. --go-grpc_out=. --go-grpc_opt=paths=source_relative stocks_broker.proto && \
	  cd ..
