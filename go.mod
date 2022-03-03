module github.com/itzmanish/micro-graphql-gateway

go 1.15

replace (
	github.com/coreos/etcd => github.com/ozonru/etcd v3.3.20-grpc1.27-origmodule+incompatible
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.8
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)

require (
	github.com/golang/protobuf v1.4.3
	github.com/graphql-go/graphql v0.7.8
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/itzmanish/go-micro/v2 v2.10.1
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
)
