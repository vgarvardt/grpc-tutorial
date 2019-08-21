NO_COLOR=\033[0m
OK_COLOR=\033[32;01m

.PHONY: deps proto

deps:
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
	@go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
	@go mod vendor

proto:
	@mkdir -p pkg/rpc
	@protoc \
		--proto_path proto \
		--go_out=plugins=grpc:pkg/rpc \
		proto/*.proto
