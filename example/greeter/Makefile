.PHONY: build

build:
	protoc \
	  -I. \
		-I../../include/graphql \
		--plugin=../../protoc-gen-graphql/protoc-gen-graphql \
	  --go_out=plugins=grpc:./greeter \
	  --micro_out=./greeter \
		--graphql_out=verbose:./greeter \
	  greeter.proto
