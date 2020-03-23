package application

import (
	"context"
	"os"

	"github.com/lukasjarosch/escrm/pkg/logger"
)

type Adapter interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

// Application represents a runnable service.
type Application struct {
	adapters []Adapter
	opts     Options
}

// New returns a new, pre-initialized application instance.
func New(options ...Option) *Application {
	opts := newOptions(options...)

	return &Application{
		opts: opts,
	}
}

func (app *Application) RegisterAdapter(adapter Adapter) {
	app.adapters = append(app.adapters, adapter)
}

func (app *Application) Run(ctx context.Context) {
	if len(app.adapters) == 0 {
		logger.Fatal(ctx, "no application adapters registered, nothing to run")
	}

	logger.Info(ctx, "starting application adapters")

	for _, adapter := range app.adapters {
		go func(a Adapter) {
			if err := a.Start(ctx); err != nil {
				logger.Fatal(ctx, "adapter start failed: %s\n", err)
			}
		}(adapter)
	}

	SignalHandler(app.shutdownFunc(ctx))
}

func (app *Application) shutdownFunc(ctx context.Context) func() {
	return func() {
		ctxWithTimeout, cancel := context.WithTimeout(ctx, app.opts.ShutdownTimeout)
		defer cancel()

		logger.Info(ctxWithTimeout, "shutting down application...\n")

		var errOccurred bool
		errChan := make(chan error, len(app.adapters))

		// stop all running adapters and stream their errors into the errChan.
		for _, adapter := range app.adapters {
			go func(adapter Adapter) {
				errChan <- adapter.Stop(ctxWithTimeout)
			}(adapter)
		}

		// for every adapter, check if there is an error
		for i := 0; i < len(app.adapters); i++ {
			if err := <-errChan; err != nil {
				errOccurred = true
				logger.Fatal(ctxWithTimeout, "adapter shutdown error: %s\n", err)
			}
		}

		if errOccurred {
			os.Exit(1)
		}

		logger.Info(ctxWithTimeout, "application was shut down gracefully")
		os.Exit(0)
	}
}
