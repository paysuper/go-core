package entrypoint

import (
	"github.com/google/wire"
	"github.com/paysuper/go-core/v2/pkg/config"
	"github.com/paysuper/go-core/v2/pkg/logger"
	"github.com/paysuper/go-core/v2/pkg/metric"
	"github.com/paysuper/go-core/v2/pkg/tracing"
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
