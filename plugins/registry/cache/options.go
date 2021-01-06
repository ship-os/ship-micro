package cache

import (
	"time"

	"github.com/ship-os/ship-micro/v2/registry/cache"
)

// WithTTL sets the cache TTL
func WithTTL(t time.Duration) cache.Option {
	return cache.WithTTL(t)
}
