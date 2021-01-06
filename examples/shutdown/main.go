package main

import (
	"log"
	"time"

	"context"
	"github.com/ship-os/ship-micro/v2"
)

func main() {
	// cancellation context
	ctx, cancel := context.WithCancel(context.Background())

	// shutdown after 5 seconds
	go func() {
		<-time.After(time.Second * 5)
		log.Println("Shutdown example: shutting down service")
		cancel()
	}()

	// create service
	service := micro.NewService(
		// with our cancellation context
		micro.Context(ctx),
	)

	// init service
	service.Init()

	// run service
	service.Run()
}
