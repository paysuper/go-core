// +build wireinject

package entrypoint

import (
	"context"
	"github.com/google/wire"
	"github.com/paysuper/go-core/v2/pkg/config"
	"github.com/paysuper/go-core/v2/pkg/invoker"
	"github.com/paysuper/go-core/v2/pkg/logger"
	"github.com/paysuper/go-core/v2/pkg/metric"
	"github.com/paysuper/go-core/v2/pkg/tracing"
)

// Build returns entrypoint instance implemented of Master interface with resolved dependencies
func Build(ctx context.Context, initial config.Initial, observer invoker.Observer) (Master, func(), error) {
	panic(wire.Build(WireSet, logger.WireSet, metric.WireSet, tracing.WireSet, config.WireSet))
}
