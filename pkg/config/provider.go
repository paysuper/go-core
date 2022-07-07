package config

import (
	"github.com/google/wire"
	"github.com/paysuper/go-core/pkg/invoker"
)

// Provider returns instance implemented Configurator interface with resolved dependencies
func Provider(initial Initial, observer invoker.Observer) (Configurator, func(), error) {
	c, e := NewProductionConfigurator(initial, observer)
	return c, func() {}, e
}

// ProviderTest returns stub/mock instance implemented Configurator interface with resolved dependencies
func ProviderTest(initial Initial, observer invoker.Observer, settings Settings) (Configurator, func(), error) {
	c, e := NewMockConfigurator(initial, observer, settings)
	return c, func() {}, e
}

var (
	WireSet     = wire.NewSet(Provider)
	WireTestSet = wire.NewSet(ProviderTest)
)
