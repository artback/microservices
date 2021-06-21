.PHONY: generate
#Path to .proto files
PROTO_PATH := api/v1/proto
GRPC_OUT := api
generate:
	 protoc -I $(PROTO_PATH) --go_out=plugins=grpc:$(GRPC_OUT) $(PROTO_PATH)/*.proto
