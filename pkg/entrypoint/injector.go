// +build wireinject

package entrypoint

import (
	"context"
	"github.com/google/wire"
	"github.com/paysuper/go-core/pkg/config"
	"github.com/paysuper/go-core/pkg/invoker"
	"github.com/paysuper/go-core/pkg/logger"
	"github.com/paysuper/go-core/pkg/metric"
	"github.com/paysuper/go-core/pkg/tracing"
)

// Build returns entrypoint instance implemented of Master interface with resolved dependencies
func Build(ctx context.Context, initial config.Initial, observer invoker.Observer) (Master, func(), error) {
	panic(wire.Build(WireSet, logger.WireSet, metric.WireSet, tracing.WireSet, config.WireSet))
}
