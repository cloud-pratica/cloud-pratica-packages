include ../Makefile

SERVICE ?= cost-provider

# proto ファイルから Go コードを生成
# e.g: make generate SERVICE=cost-provider
generate-proto: .check-service
	@protoc \
		--proto_path=./$(SERVICE)/api/v1/ \
		--go_out=./$(SERVICE)/gen \
		--go_opt=paths=source_relative \
		--go-grpc_out=./$(SERVICE)/gen \
		--go-grpc_opt=paths=source_relative \
		service.proto

.PHONY: generate-proto

.check-service:
ifndef SERVICE
	$(error SERVICE is required.)
endif
