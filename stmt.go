package wrapper

import (
	"context"
	"database/sql/driver"
	"fmt"
	"time"
)

// wrapperStmt defines a wrapper for driver.Stmt
type wrapperStmt struct {
	stmt driver.Stmt
	opts Options
}

// Close implements driver.Stmt Close
func (w *wrapperStmt) Close() error {
	labels := []string{labelMethod, "Close"}
	ts := time.Now()
	err := w.stmt.Close()
	te := time.Since(ts).Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
	} else {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	}
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
	return err
}

// NumInput implements driver.Stmt NumInput
func (w *wrapperStmt) NumInput() int {
	return w.stmt.NumInput()
}

// Exec implements driver.Stmt Exec
func (w *wrapperStmt) Exec(args []driver.Value) (driver.Result, error) {
	// nolint:staticcheck
	labels := []string{labelMethod, "Exec"}
	ts := time.Now()
	res, err := w.stmt.Exec(args)
	te := time.Since(ts).Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
	} else {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	}
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
	return res, err
}

// Query implements driver.Stmt Query
func (w *wrapperStmt) Query(args []driver.Value) (driver.Rows, error) {
	labels := []string{labelMethod, "Query"}
	ts := time.Now()
	// nolint:staticcheck
	rows, err := w.stmt.Query(args)
	te := time.Since(ts).Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
	} else {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	}
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
	return rows, err
}

// ExecContext implements driver.ExecerContext ExecContext
func (w *wrapperStmt) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	nctx, span := w.opts.Tracer.Start(ctx, "ExecContext")
	span.AddLabels("method", "ExecContext")
	name := getQueryName(ctx)
	if name != "" {
		span.AddLabels("query", name)
	} else {
		name = labelUnknown
	}
	defer span.Finish()
	if len(args) > 0 {
		span.AddLabels("args", fmt.Sprintf("%v", namedValueToLabels(args)))
	}
	labels := []string{labelMethod, "ExecContext", labelQuery, name}
	if execerContext, ok := w.stmt.(driver.ExecerContext); ok {
		ts := time.Now()
		res, err := execerContext.ExecContext(nctx, query, args)
		te := time.Since(ts).Seconds()
		if err != nil {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
			span.AddLabels("error", true)
			span.AddLabels("err", err.Error())
		} else {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
		}
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
		return res, err
	}
	values, err := namedValueToValue(args)
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
		return nil, err
	}
	// nolint:staticcheck
	res, err := w.Exec(values)
	if err != nil {
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
	}
	return res, err
}

// QueryContext implements Driver.QueryerContext QueryContext
func (w *wrapperStmt) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	nctx, span := w.opts.Tracer.Start(ctx, "QueryContext")
	span.AddLabels("method", "QueryContext")
	name := getQueryName(ctx)
	if name != "" {
		span.AddLabels("query", name)
	} else {
		name = labelUnknown
	}
	defer span.Finish()
	if len(args) > 0 {
		span.AddLabels("args", fmt.Sprintf("%v", namedValueToLabels(args)))
	}
	if queryerContext, ok := w.stmt.(driver.QueryerContext); ok {
		rows, err := queryerContext.QueryContext(nctx, query, args)
		if err != nil {
			span.AddLabels("error", true)
			span.AddLabels("err", err.Error())
		}
		return rows, err
	}
	values, err := namedValueToValue(args)
	if err != nil {
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
		return nil, err
	}
	// nolint:staticcheck
	rows, err := w.Query(values)
	if err != nil {
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
	}
	return rows, err
}
