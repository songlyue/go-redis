package tcp

import (
	"context"
	"net"
)

// Handler represents application server over tcp
type Handler interface {
	Handler(ctx context.Context, conn net.Conn)
	Close() error
}
