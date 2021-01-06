// Package service provides the broker service client
package service

import (
	"github.com/ship-os/ship-micro/v2/broker"
	"github.com/ship-os/ship-micro/v2/broker/service"
	"github.com/ship-os/ship-micro/v2/cmd"
)

func init() {
	cmd.DefaultBrokers["service"] = NewBroker
}

func NewBroker(opts ...broker.Option) broker.Broker {
	return service.NewBroker(opts...)
}
