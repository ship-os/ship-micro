package main

import (
	"context"
	"log"

	hello "github.com/ship-os/ship-micro/examples/greeter/srv/proto/hello"
	k8s "github.com/ship-os/ship-micro/examples/kubernetes/go/micro"
	"github.com/ship-os/ship-micro/v2"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Print("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := k8s.NewService(
		micro.Name("greeter"),
	)

	service.Init()

	hello.RegisterSayHandler(service.Server(), new(Say))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
