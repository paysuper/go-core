package metric

import (
	"context"
	"github.com/paysuper/go-core/pkg/invoker"
	"github.com/uber-go/tally"
	promreporter "github.com/uber-go/tally/prometheus"
	tallystatsd "github.com/uber-go/tally/statsd"
	"time"
)

const (
	Prefix       = "go-core.metric"
	UnmarshalKey = "metric"
)

// StatsDCfg is a setting for tally statsd client
type StatsDCfg struct {
	Addr, Prefix  string
	FlushInterval time.Duration
	FlushBytes    int
	Options       tallystatsd.Options
}

// PrometheusCfg is a setting for tally prometheus connector
type PrometheusCfg struct {
	Address string `default:"http://0.0.0.0:9090/metrics"`
	Options promreporter.Options
}

// Config is a general metric config settings
type Config struct {
	Enabled    bool
	StatsD     StatsDCfg
	Prometheus PrometheusCfg
	Scope      tally.ScopeOptions
	Interval   time.Duration
	invoker    *invoker.Invoker
}

// OnReload
func (c *Config) OnReload(callback func(ctx context.Context)) {
	c.invoker.OnReload(callback)
}

// Reload
func (c *Config) Reload(ctx context.Context) {
	c.invoker.Reload(ctx)
}

type (
	Scope        = tally.Scope
	Counter      = tally.Counter
	Gauge        = tally.Gauge
	Timer        = tally.Timer
	Histogram    = tally.Histogram
	Buckets      = tally.Buckets
	Capabilities = tally.Capabilities
	Stopwatch    = tally.Stopwatch
)
