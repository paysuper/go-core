// +build wireinject

package config

import (
	"github.com/google/wire"
	"github.com/paysuper/go-core/pkg/invoker"
)

// Build returns instance implemented Configurator interface with resolved dependencies
func Build(initial Initial, observer invoker.Observer) (Configurator, func(), error) {
	panic(wire.Build(Provider))
}

// BuildTest returns stub/mock instance implemented Configurator interface with resolved dependencies
func BuildTest(initial Initial, observer invoker.Observer, settings Settings) (Configurator, func(), error) {
	panic(wire.Build(WireTestSet))
}
