package app

import (
	"context"
	"github.com/go-nunu/nunu-layout-advanced/pkg/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	name    string
	servers []server.Server
}

type Option func(a *App)

func NewApp(opts ...Option) *App {
	a := &App{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func WithServer(servers ...server.Server) Option {
	return func(a *App) {
		a.servers = servers
	}
}
func WithName(name string) Option {
	return func(a *App) {
		a.name = name
	}
}

func (a *App) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	endpoints := make([]string, 0)
	for _, srv := range a.servers {
		if s, ok := srv.(server.Endpointer); ok {
			e, err := s.Endpoint()
			if err != nil {
				return err
			}
			endpoints = append(endpoints, e.String())
		}
		go func(srv server.Server) {
			err := srv.Start(ctx)
			if err != nil {
				cancel()
				log.Fatal("app start err:", err)
			}
		}(srv)
	}

	<-signals
	cancel()
	for _, srv := range a.servers {
		err := srv.Stop(ctx)
		if err != nil {
			log.Fatal("app stop err:", err)
		}
	}

	return nil
}
