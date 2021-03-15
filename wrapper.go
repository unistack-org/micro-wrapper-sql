package wrapper

import (
	"context"
	"database/sql"
	"time"

	"github.com/unistack-org/micro/v3/logger"
	"github.com/unistack-org/micro/v3/meter"
	"github.com/unistack-org/micro/v3/tracer"
)

var (
	DefaultStatsInterval = 5 * time.Second
	// default metric prefix
	DefaultMetricPrefix = "micro_sql_"
	// default label prefix
	DefaultLabelPrefix = "micro_"
)

var (
	DatabaseHostLabel   = "database_host"
	DatabaseNameLabel   = "database_name"
	ServiceNameLabel    = "service_name"
	ServiceVersionLabel = "service_version"
	ServiceIDLabel      = "service_id"

	MaxOpenConnectionsLabel = "max_open_connections"
	OpenConnectionsLabel    = "open_connections"
	InuseConnectionsLabel   = "inuse_connections"
	IdleConnectionsLabel    = "idle_connections"
	WaitConnectionsLabel    = "wait_connections"
	BlockedSecondsLabel     = "blocked_seconds"
	MaxIdleClosedLabel      = "max_idle_closed"
	MaxLifetimeClosedLabel  = "max_lifetime_closed"

	//srequest_total // counter
	//slatency_microseconds // summary
	//srequest_duration_seconds // histogramm

)

type Options struct {
	Logger         logger.Logger
	Meter          meter.Meter
	Tracer         tracer.Tracer
	DatabaseHost   string
	DatabaseName   string
	ServiceName    string
	ServiceVersion string
	ServiceID      string
}

type Option func(*Options)

func DatabaseHost(host string) Option {
	return func(opts *Options) {
		opts.DatabaseHost = host
	}
}

func DatabaseName(name string) Option {
	return func(opts *Options) {
		opts.DatabaseName = name
	}
}

func ServiceName(name string) Option {
	return func(opts *Options) {
		opts.ServiceName = name
	}
}

func ServiceVersion(version string) Option {
	return func(opts *Options) {
		opts.ServiceVersion = version
	}
}

func ServiceID(id string) Option {
	return func(opts *Options) {
		opts.ServiceID = id
	}
}

func Meter(m meter.Meter) Option {
	return func(opts *Options) {
		opts.Meter = m
	}
}
func Logger(l logger.Logger) Option {
	return func(opts *Options) {
		opts.Logger = l
	}
}

func Tracer(t tracer.Tracer) Option {
	return func(opts *Options) {
		opts.Tracer = t
	}
}

type queryKey struct{}

func QueryName(ctx context.Context, name string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithValue(ctx, queryKey{}, name)
}

func getName(ctx context.Context) string {
	name := "Unknown"

	val, ok := ctx.Value(queryKey{}).(string)
	if ok && len(val) > 0 {
		name = val
	}

	return name
}

type Wrapper struct {
	db   *sql.DB
	opts Options
}

func (w *Wrapper) collect() {
	labels := []string{
		DatabaseHostLabel, w.opts.DatabaseHost,
		DatabaseNameLabel, w.opts.DatabaseName,
		ServiceNameLabel, w.opts.ServiceName,
		ServiceVersionLabel, w.opts.ServiceVersion,
		ServiceIDLabel, w.opts.ServiceID,
	}

	ticker := time.NewTicker(DefaultStatsInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if w.db == nil {
				continue
			}

			stats := w.db.Stats()
			w.opts.Meter.FloatCounter(MaxOpenConnectionsLabel, meter.Labels(labels...)).Set(float64(stats.MaxOpenConnections))
			w.opts.Meter.FloatCounter(OpenConnectionsLabel, meter.Labels(labels...)).Set(float64(stats.OpenConnections))
			w.opts.Meter.FloatCounter(InuseConnectionsLabel, meter.Labels(labels...)).Set(float64(stats.InUse))
			w.opts.Meter.FloatCounter(IdleConnectionsLabel, meter.Labels(labels...)).Set(float64(stats.Idle))
			w.opts.Meter.FloatCounter(WaitConnectionsLabel, meter.Labels(labels...)).Set(float64(stats.WaitCount))
			w.opts.Meter.FloatCounter(BlockedSecondsLabel, meter.Labels(labels...)).Set(stats.WaitDuration.Seconds())
			w.opts.Meter.FloatCounter(MaxIdleClosedLabel, meter.Labels(labels...)).Set(float64(stats.MaxIdleClosed))
			w.opts.Meter.FloatCounter(MaxLifetimeClosedLabel, meter.Labels(labels...)).Set(float64(stats.MaxLifetimeClosed))
		}
	}
}
