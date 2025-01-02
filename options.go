package wrapper

import (
	"context"
	"time"

	"go.unistack.org/micro/v3/logger"
	"go.unistack.org/micro/v3/meter"
	"go.unistack.org/micro/v3/tracer"
)

// DefaultMeterStatsInterval holds default stats interval
var DefaultMeterStatsInterval = 5 * time.Second

var (
	MaxOpenConnections = "micro_sql_max_open_conn"
	OpenConnections    = "micro_sql_open_conn"
	InuseConnections   = "micro_sql_inuse_conn"
	IdleConnections    = "micro_sql_idle_conn"
	WaitConnections    = "micro_sql_waited_conn"
	BlockedSeconds     = "micro_sql_blocked_seconds"
	MaxIdleClosed      = "micro_sql_max_idle_closed"
	MaxIdletimeClosed  = "micro_sql_closed_max_idle"
	MaxLifetimeClosed  = "micro_sql_closed_max_lifetime"

	meterRequestTotal               = "micro_sql_request_total"
	meterRequestLatencyMicroseconds = "micro_sql_latency_microseconds"
	meterRequestDurationSeconds     = "micro_sql_request_duration_seconds"

	labelUnknown  = "unknown"
	labelQuery    = "db_statement"
	labelMethod   = "db_method"
	labelStatus   = "status"
	labelSuccess  = "success"
	labelFailure  = "failure"
	labelHost     = "db_host"
	labelDatabase = "db_name"
)

// Options struct holds wrapper options
type Options struct {
	Meter              meter.Meter
	Tracer             tracer.Tracer
	DatabaseHost       string
	DatabaseName       string
	MeterStatsInterval time.Duration
	LoggerLevel        logger.Level
	LoggerEnabled      bool
	LoggerObserver     func(ctx context.Context, method string, name string, td time.Duration, err error) []interface{}
}

// Option func signature
type Option func(*Options)

// NewOptions create new Options struct from provided option slice
func NewOptions(opts ...Option) Options {
	options := Options{
		Meter:              meter.DefaultMeter,
		Tracer:             tracer.DefaultTracer,
		MeterStatsInterval: DefaultMeterStatsInterval,
	}
	for _, o := range opts {
		o(&options)
	}

	options.Meter = options.Meter.Clone(
		meter.Labels(
			labelHost, options.DatabaseHost,
			labelDatabase, options.DatabaseName,
		),
	)

	return options
}

// MetricInterval specifies stats interval for *sql.DB
func MetricInterval(td time.Duration) Option {
	return func(o *Options) {
		o.MeterStatsInterval = td
	}
}

func DatabaseHost(host string) Option {
	return func(o *Options) {
		o.DatabaseHost = host
	}
}

func DatabaseName(name string) Option {
	return func(o *Options) {
		o.DatabaseName = name
	}
}

// Meter passes meter.Meter to wrapper
func Meter(m meter.Meter) Option {
	return func(o *Options) {
		o.Meter = m
	}
}

// Tracer passes tracer.Tracer to wrapper
func Tracer(t tracer.Tracer) Option {
	return func(o *Options) {
		o.Tracer = t
	}
}

type queryNameKey struct{}

// QueryName passes query name to wrapper func
func QueryName(ctx context.Context, name string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, queryNameKey{}, name)
}

func getQueryName(ctx context.Context) string {
	if v, ok := ctx.Value(queryNameKey{}).(string); ok && v != labelUnknown {
		return v
	}
	return getCallerName()
}
