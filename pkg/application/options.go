package application

import (
	"time"
)

type Options struct {
	ShutdownTimeout time.Duration
}

func ShutdownTimeout(timeout time.Duration) Option {
    return func(opts *Options) {
        opts.ShutdownTimeout = timeout
    }
}

func newOptions(opts ...Option) Options {
	opt := Options{
		ShutdownTimeout: 5 * time.Second,
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

type Option func(*Options)