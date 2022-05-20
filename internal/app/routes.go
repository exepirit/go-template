package app

import (
	"github.com/exepirit/go-template/internal/service/greeter"
	"github.com/exepirit/go-template/pkg/routing"
)

func NewRoutes(
	greeter *greeter.Endpoints,
) routing.Bindable {
	return routing.Route("/",
		routing.Route("/api/v1",
			routing.Route("/greeting", greeter),
		),
	)
}
