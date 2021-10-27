package wrapper // import "go.unistack.org/micro-wrapper-sql/v3"

import (
	"context"
	"database/sql"
	"time"

	"go.unistack.org/micro/v3/meter"
)

type Statser interface {
	Stats() sql.DBStats
}

func NewStatsMeter(ctx context.Context, db Statser, opts ...Option) {
	options := NewOptions(opts...)

	m := options.Meter.Clone(
		meter.LabelPrefix(options.MeterLabelPrefix),
		meter.MetricPrefix(options.MeterMetricPrefix),
		meter.Labels(
			labelHost, options.DatabaseHost,
			labelDatabase, options.DatabaseName,
		),
	)

	go func() {
		ticker := time.NewTicker(options.MeterStatsInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				if db == nil {
					return
				}
				stats := db.Stats()
				m.Counter(MaxOpenConnections).Set(uint64(stats.MaxOpenConnections))
				m.Counter(OpenConnections).Set(uint64(stats.OpenConnections))
				m.Counter(InuseConnections).Set(uint64(stats.InUse))
				m.Counter(IdleConnections).Set(uint64(stats.Idle))
				m.Counter(WaitConnections).Set(uint64(stats.WaitCount))
				m.FloatCounter(BlockedSeconds).Set(stats.WaitDuration.Seconds())
				m.Counter(MaxIdleClosed).Set(uint64(stats.MaxIdleClosed))
				m.Counter(MaxLifetimeClosed).Set(uint64(stats.MaxLifetimeClosed))
			}
		}
	}()
}
