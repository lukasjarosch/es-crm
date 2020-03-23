package main

import (
	"context"

	"github.com/lukasjarosch/escrm/pkg/application"
	"github.com/lukasjarosch/escrm/pkg/logger"
)

func init() {
	logger.NewGlobalLogger(logger.DefaultLevel)
}

func main() {
	logger.Info(nil, "starting example v{git-commit} ({build-date})")
	app := application.New()

	ta := &TestAdapter{}

	app.RegisterAdapter(ta)


	app.Run(context.Background())
}

type TestAdapter struct {
}

func (a *TestAdapter) Start(ctx context.Context) error {
	logger.Info(ctx, "starting TestAdapter")
	return nil
}

func (a *TestAdapter) Stop(ctx context.Context) error {
	logger.Info(ctx, "stopping TestAdapter")
	return nil
}
