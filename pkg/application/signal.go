package application

import (
	"os"
	"os/signal"
	"syscall"
)

// shutdownSignals are all signals which the SignalHandler is notified of.
var shutdownSignals = []os.Signal{
	syscall.SIGHUP,
	syscall.SIGINT,
	syscall.SIGTERM,
	syscall.SIGQUIT,
}

// SignalHandler, once called will start listening for one of the defined shutdownSignals.
// If a signal is caught, the exit-code will be determined by whether the signal is an expected shutdownSignal (exit 0)
// or an unexpected interrupt (exit 1)
//
// Before os.Exit is called, the stopFunction is executed.
func SignalHandler(stopFunction func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, shutdownSignals...)

	exitCodeChan := make(chan int)
	go func() {
		for {
			sig := <-sigChan

			for _, shutdownSignal := range shutdownSignals {
				if sig.String() == shutdownSignal.String() {
					exitCodeChan <- 0
				} else {
					exitCodeChan <- 1
				}
			}
		}
	}()

	exitCode := <-exitCodeChan
	stopFunction()
	os.Exit(exitCode)
}