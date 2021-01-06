package subscriber

import (
	"context"
	"github.com/ship-os/ship-micro/v2/util/log"

	example "github.com/ship-os/ship-micro/examples/template/srv/proto/example"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *example.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
