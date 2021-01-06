// Package registry uses the go-micro registry for selection
package registry

import (
	"github.com/ship-os/ship-micro/v2/cmd"
	"github.com/ship-os/ship-micro/v2/selector"
)

func init() {
	cmd.DefaultSelectors["registry"] = NewSelector
}

// NewSelector returns a new registry selector
func NewSelector(opts ...selector.Option) selector.Selector {
	return selector.NewSelector(opts...)
}
