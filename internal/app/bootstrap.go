package app

import (
	"context"

	"github.com/exepirit/go-template/internal/config"
	"github.com/exepirit/go-template/internal/infrastructure"
	"github.com/exepirit/go-template/pkg/routing"
	"go.uber.org/fx"
)

func bootstrap(
	lifecycle fx.Lifecycle,
	cfg config.Config,
	server infrastructure.Server,
	routes routing.Bindable,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			server.Bind(routes)

			go func(server infrastructure.Server) {
				if err := server.ListenAndServe(); err != nil {
					panic(err)
				}
			}(server)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
}
