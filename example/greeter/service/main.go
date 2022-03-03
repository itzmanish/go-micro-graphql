package main

import (
	"context"
	"log"

	"github.com/itzmanish/go-micro/v2"
	"github.com/itzmanish/go-micro/v2/server/grpc"
	"github.com/itzmanish/micro-graphql-gateway/example/greeter/greeter"
)

func main() {
	service := micro.NewService(
		micro.Name("itzmanish.greeter"),
		micro.Server(grpc.NewServer()),
	)

	service.Init()

	err := greeter.RegisterGreeterHandler(service.Server(), new(greeterHandler))
	if err != nil {
		log.Fatal(err)
	}
	service.Run()
}

type greeterHandler struct{}

func (gh *greeterHandler) SayHello(ctx context.Context, in *greeter.HelloRequest, out *greeter.HelloReply) error {
	out.Message = "Hola " + in.GetName()
	return nil
}
func (gh *greeterHandler) SayGoodbye(ctx context.Context, in *greeter.GoodbyeRequest, out *greeter.GoodbyeReply) error {
	out.Message = "Goodbye " + in.GetName()
	return nil
}
