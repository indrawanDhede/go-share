//go:build wireinject
// +build wireinject

package testing

import "github.com/google/wire"

func InitializeRoute() *SimpleRoute {
	wire.Build(
		NewSimpleRepository, NewSimpleService, NewSimpleController, NewSimpleRoute,
	)
	return nil
}
