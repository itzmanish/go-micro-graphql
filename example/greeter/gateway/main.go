package main

import (
	"log"
	"net/http"

	"github.com/itzmanish/go-micro/v2/client/grpc"
	"github.com/itzmanish/micro-graphql-gateway/example/greeter/greeter"
	"github.com/itzmanish/micro-graphql-gateway/runtime"
)

func main() {
	mux := runtime.NewServeMux()
	if err := greeter.RegisterGreeterGraphqlHandler(mux, "itzmanish.greeter", grpc.NewClient()); err != nil {
		log.Fatalln(err)
	}
	http.Handle("/graphql", mux)
	log.Println(http.ListenAndServe(":8080", nil))

}
