package main

import (
	"github.com/ship-os/ship-micro/examples/template/fnc/handler"
	"github.com/ship-os/ship-micro/examples/template/fnc/subscriber"
	"github.com/ship-os/ship-micro/v2"
	"github.com/ship-os/ship-micro/v2/util/log"
)

func main() {
	// New Service
	function := micro.NewFunction(
		micro.Name("go.micro.fnc.template"),
		micro.Version("latest"),
	)

	// Register Handler
	function.Handle(new(handler.Example))

	// Register Struct as Subscriber
	function.Subscribe("go.micro.fnc.template", new(subscriber.Example))

	// Initialise function
	function.Init()

	// Run service
	if err := function.Run(); err != nil {
		log.Fatal(err)
	}
}
