package subscriber

import (
	"context"

	"github.com/ship-os/ship-micro/v2/util/log"

	example "github.com/ship-os/ship-micro/examples/template/fnc/proto/example"
)

type Example struct{}

func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}
