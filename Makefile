.PHONY: generate
#Path to .proto files
PROTO_PATH := api/v1/proto
GRPC_OUT := api
generate:
	 protoc -I $(PROTO_PATH)  --go_out=plugins=grpc:$(GRPC_OUT) $(PROTO_PATH)/pds.proto
	 protoc -I $(PROTO_PATH) -I  $(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis --go_out=plugins=grpc:$(GRPC_OUT)  --grpc-gateway_out=logtostderr=true:$(GRPC_OUT)  $(PROTO_PATH)/client.proto
