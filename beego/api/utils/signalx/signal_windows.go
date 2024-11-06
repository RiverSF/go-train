//go:build windows
// +build windows

package signalx

import (
	"os"
	"os/signal"
	"syscall"
)

func SignalHandler(fc func()) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			signal.Stop(ch)
			fc()
		}
	}
}
