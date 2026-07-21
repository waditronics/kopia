//go:build windows
// +build windows

package apiclient

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
)

func dialNamedPipe(ctx context.Context, pipeName string) (net.Conn, error) {
	return winio.DialPipeContext(ctx, pipeName)
}
