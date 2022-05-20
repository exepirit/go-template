package app

import (
	"github.com/exepirit/go-template/internal/config"
	"github.com/exepirit/go-template/internal/infrastructure"
	"github.com/exepirit/go-template/internal/service"
	"go.uber.org/fx"
)

var Module = fx.Options(
	config.Module,
	infrastructure.Module,
	service.Module,
	fx.Provide(NewRoutes),
	fx.Invoke(bootstrap),
)
