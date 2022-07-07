package provider

import (
	"github.com/google/wire"
	"github.com/paysuper/go-core/pkg/config"
	"github.com/paysuper/go-core/pkg/logger"
	"github.com/paysuper/go-core/pkg/metric"
	"github.com/paysuper/go-core/pkg/tracing"
)

var Set = wire.NewSet(
	logger.WireSet,
	config.WireSet,
	metric.WireSet,
	tracing.WireSet,
)
