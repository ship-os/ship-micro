package main

import (
	"github.com/ship-os/ship-micro/v2/util/log"

	"github.com/ship-os/ship-micro/examples/template/api/client"
	"github.com/ship-os/ship-micro/examples/template/api/handler"
	"github.com/ship-os/ship-micro/v2"

	example "github.com/ship-os/ship-micro/examples/template/api/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.template"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Initialise service
	service.Init(
		// create wrap for the Example srv client
		micro.WrapHandler(client.ExampleWrapper(service)),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
