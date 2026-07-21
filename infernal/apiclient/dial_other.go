//go:build !windows
// +build !windows

package apiclient

import (
	"context"
	"net"
)

func dialNamedPipe(ctx context.Context, pipeName string) (net.Conn, error) {
	return nil, nil
}
