package grpc

import (
	"github.com/ship-os/ship-micro/v2/config/source"
	proto "github.com/ship-os/ship-micro/plugins/config/source/grpc/v2/proto"
)

type watcher struct {
	stream proto.Source_WatchClient
}

func newWatcher(stream proto.Source_WatchClient) (*watcher, error) {
	return &watcher{
		stream: stream,
	}, nil
}

func (w *watcher) Next() (*source.ChangeSet, error) {
	rsp, err := w.stream.Recv()
	if err != nil {
		return nil, err
	}
	return toChangeSet(rsp.ChangeSet), nil
}

func (w *watcher) Stop() error {
	return w.stream.CloseSend()
}
