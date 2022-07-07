package provider

import (
	"github.com/google/wire"
	"github.com/paysuper/go-core/v2/pkg/config"
	"github.com/paysuper/go-core/v2/pkg/logger"
	"github.com/paysuper/go-core/v2/pkg/metric"
	"github.com/paysuper/go-core/v2/pkg/tracing"
)

var Set = wire.NewSet(
	logger.WireSet,
	config.WireSet,
	metric.WireSet,
	tracing.WireSet,
)
