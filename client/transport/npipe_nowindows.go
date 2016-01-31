// +build -windows

package transport

import (
	"net"
	"syscall"
	"time"
)

// DialPipe connects to a Windows named pipe. This is not supported on other OSes.
func DialPipe(_ addr, _ time.Duration) (net.Conn, error) {
	return nil, syscall.EAFNOSUPPORT
}
