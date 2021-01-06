package hystrix

import (
	"testing"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/ship-os/ship-micro/v2/client"
	"github.com/ship-os/ship-micro/v2/selector"
	"github.com/ship-os/ship-micro/v2/registry/memory"

	"context"
)

func TestBreaker(t *testing.T) {
	// setup
	r := memory.NewRegistry()
	s := selector.NewSelector(selector.Registry(r))

	c := client.NewClient(
		// set the selector
		client.Selector(s),
		// add the breaker wrapper
		client.Wrap(NewClientWrapper()),
	)

	req := c.NewRequest("test.service", "Test.Method", map[string]string{
		"foo": "bar",
	}, client.WithContentType("application/json"))

	var rsp map[string]interface{}

	// Force to point of trip
	for i := 0; i < (hystrix.DefaultVolumeThreshold * 3); i++ {
		c.Call(context.TODO(), req, rsp)
	}

	err := c.Call(context.TODO(), req, rsp)
	if err == nil {
		t.Error("Expecting tripped breaker, got nil error")
	}

	if err.Error() != "hystrix: circuit open" {
		t.Errorf("Expecting tripped breaker, got %v", err)
	}
}
