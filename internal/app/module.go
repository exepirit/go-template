package app

import (
	"github.com/exepirit/go-template/internal/config"
	"github.com/exepirit/go-template/internal/infrastructure"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
	infrastructure.Module,
	fx.Provide(NewRoutes),
	fx.Invoke(bootstrap),
)
