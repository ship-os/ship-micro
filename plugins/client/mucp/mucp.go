// Package mucp provides an mucp client
package mucp

import (
	"github.com/ship-os/ship-micro/v2/cmd"
	"github.com/ship-os/ship-micro/v2/client"
)

func init() {
	cmd.DefaultClients["mucp"] = NewClient
}

// NewClient returns a new micro client interface
func NewClient(opts ...client.Option) client.Client {
	return client.NewClient(opts...)
}
