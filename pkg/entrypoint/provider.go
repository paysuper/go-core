package entrypoint

import (
	"github.com/google/wire"
	"github.com/paysuper/go-core/pkg/config"
	"github.com/paysuper/go-core/pkg/logger"
	"github.com/paysuper/go-core/pkg/metric"
	"github.com/paysuper/go-core/pkg/tracing"
)

type AppSet struct {
	Logger logger.Logger
	Metric metric.Scope
	Tracer tracing.Tracer
}

// Provider returns entrypoint instance implemented of Master interface with resolved dependencies
func Provider(set AppSet, initial config.Initial) (Master, func(), error) {
	e, r := NewEntryPoint(set, initial)
	return e, func() {}, r
}

var (
	WireSet = wire.NewSet(Provider, wire.Struct(new(AppSet), "*"))
)
