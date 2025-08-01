package application

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	ctx        context.Context
	cancel     context.CancelFunc
	log        *slog.Logger
	cleanup    []func(context.Context) error
	shutdownCh chan os.Signal
}

func NewApp(parent context.Context, log *slog.Logger) *App {
	newCtx, cancel := context.WithCancel(parent)
	return &App{
		ctx:        newCtx,
		cancel:     cancel,
		log:        log,
		shutdownCh: make(chan os.Signal),
	}
}

func (a *App) Run() {
	const op = "application.Start"

	log := a.log.With(slog.String("op", op))

	log.Info("starting application")

	signal.Notify(a.shutdownCh, os.Interrupt, syscall.SIGTERM)

	<-a.shutdownCh
}

func (a *App) Shutdown() {
	const op = "application.Shutdown"

	log := a.log.With(slog.String("op", op))

	log.Info("shutting down application...")

	a.cancel()

	for i := len(a.cleanup); i >= 0; i-- {
		if err := a.cleanup[i](a.ctx); err != nil {
			log.Error(err.Error())
		}
	}

	log.Info("shutdown application completed")
}

func (a *App) RegisterCleanup(f func(context.Context) error) {
	a.cleanup = append(a.cleanup, f)
}
