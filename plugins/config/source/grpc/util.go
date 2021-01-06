package grpc

import (
	"time"

	"github.com/ship-os/ship-micro/v2/config/source"
	proto "github.com/ship-os/ship-micro/plugins/config/source/grpc/v2/proto"
)

func toChangeSet(c *proto.ChangeSet) *source.ChangeSet {
	return &source.ChangeSet{
		Data:      c.Data,
		Checksum:  c.Checksum,
		Format:    c.Format,
		Timestamp: time.Unix(c.Timestamp, 0),
		Source:    c.Source,
	}
}
