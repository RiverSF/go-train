//go:build linux
// +build linux

package signalx

import (
	"os"
	"os/signal"
	"syscall"
)

func SignalHandler(fc func()) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2:
			signal.Stop(ch)
			fc()
		}
	}
}
