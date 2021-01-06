package main

import (
	"fmt"

	"github.com/ship-os/ship-micro/v2"
	"github.com/ship-os/ship-micro/v2/server"
)

func main() {
	service := micro.NewService()

	service.Server().Init(
		server.Wait(nil),
	)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
