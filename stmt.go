package wrapper

import (
	"context"
	"database/sql/driver"
	"fmt"
	"time"

	"go.unistack.org/micro/v3/tracer"
)

var (
	_ driver.Stmt              = (*wrapperStmt)(nil)
	_ driver.StmtQueryContext  = (*wrapperStmt)(nil)
	_ driver.StmtExecContext   = (*wrapperStmt)(nil)
	_ driver.NamedValueChecker = (*wrapperStmt)(nil)
)

// wrapperStmt defines a wrapper for driver.Stmt
type wrapperStmt struct {
	stmt  driver.Stmt
	opts  Options
	query string
	ctx   context.Context
}

// Close implements driver.Stmt Close
func (w *wrapperStmt) Close() error {
	var ctx context.Context
	if w.ctx != nil {
		ctx = w.ctx
	} else {
		ctx = context.Background()
	}
	labels := []string{labelMethod, "Close"}
	ts := time.Now()
	err := w.stmt.Close()
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
	} else {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	}
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "Close", getCallerName(), td, err)).Log(ctx, w.opts.LoggerLevel)
	}
	return err
}

// NumInput implements driver.Stmt NumInput
func (w *wrapperStmt) NumInput() int {
	return w.stmt.NumInput()
}

// CheckNamedValue implements driver.NamedValueChecker
func (w *wrapperStmt) CheckNamedValue(v *driver.NamedValue) error {
	s, ok := w.stmt.(driver.NamedValueChecker)
	if !ok {
		return driver.ErrSkip
	}
	return s.CheckNamedValue(v)
}

// Exec implements driver.Stmt Exec
func (w *wrapperStmt) Exec(args []driver.Value) (driver.Result, error) {
	var ctx context.Context
	if w.ctx != nil {
		ctx = w.ctx
	} else {
		ctx = context.Background()
	}
	labels := []string{labelMethod, "Exec"}
	ts := time.Now()
	res, err := w.stmt.Exec(args) // nolint:staticcheck
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
	} else {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	}
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "Exec", getCallerName(), td, err)).Log(ctx, w.opts.LoggerLevel)
	}

	return res, err
}

// Query implements driver.Stmt Query
func (w *wrapperStmt) Query(args []driver.Value) (driver.Rows, error) {
	var ctx context.Context
	if w.ctx != nil {
		ctx = w.ctx
	} else {
		ctx = context.Background()
	}
	labels := []string{labelMethod, "Query"}
	ts := time.Now()
	rows, err := w.stmt.Query(args) // nolint:staticcheck
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
	} else {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	}
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "Query", getCallerName(), td, err)).Log(ctx, w.opts.LoggerLevel)
	}

	return rows, err
}

// ColumnConverter implements driver.ColumnConverter
func (w *wrapperStmt) ColumnConverter(idx int) driver.ValueConverter {
	s, ok := w.stmt.(driver.ColumnConverter) // nolint:staticcheck
	if !ok {
		return nil
	}
	return s.ColumnConverter(idx)
}

// ExecContext implements driver.StmtExecContext ExecContext
func (w *wrapperStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	var nctx context.Context
	var span tracer.Span

	if w.ctx != nil {
		nctx, span = w.opts.Tracer.Start(w.ctx, "sdk.database", tracer.WithSpanKind(tracer.SpanKindClient))
	} else {
		nctx, span = w.opts.Tracer.Start(ctx, "sdk.database", tracer.WithSpanKind(tracer.SpanKindClient))
	}
	span.AddLabels("method", "ExecContext")
	name := getQueryName(ctx)
	if name != "" {
		span.AddLabels("db.query", name)
	} else {
		name = getCallerName()
	}
	defer span.Finish()
	if len(args) > 0 {
		span.AddLabels("db.args", fmt.Sprintf("%v", namedValueToLabels(args)))
	}
	labels := []string{labelMethod, "ExecContext", labelQuery, name}

	if conn, ok := w.stmt.(driver.StmtExecContext); ok {
		ts := time.Now()
		res, err := conn.ExecContext(nctx, args)
		td := time.Since(ts)
		te := td.Seconds()
		if err != nil {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
			span.SetStatus(tracer.SpanStatusError, err.Error())
		} else {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
		}
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

		if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
			w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "ExecContext", name, td, err)).Log(ctx, w.opts.LoggerLevel)
		}
		return res, err
	}

	values, err := namedValueToValue(args)
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		span.SetStatus(tracer.SpanStatusError, err.Error())

		if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
			w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "ExecContext", name, 0, err)).Log(ctx, w.opts.LoggerLevel)
		}
		return nil, err
	}
	ts := time.Now()
	res, err := w.Exec(values) // nolint:staticcheck
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		span.SetStatus(tracer.SpanStatusError, err.Error())
	} else {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	}

	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "ExecContext", name, td, err)).Log(ctx, w.opts.LoggerLevel)
	}

	return res, err
}

// QueryContext implements driver.StmtQueryContext StmtQueryContext
func (w *wrapperStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	var nctx context.Context
	var span tracer.Span

	if w.ctx != nil {
		nctx, span = w.opts.Tracer.Start(w.ctx, "sdk.database", tracer.WithSpanKind(tracer.SpanKindClient))
	} else {
		nctx, span = w.opts.Tracer.Start(ctx, "sdk.database", tracer.WithSpanKind(tracer.SpanKindClient))
	}
	span.AddLabels("method", "QueryContext")
	name := getQueryName(ctx)
	if name != "" {
		span.AddLabels("db.query", name)
	} else {
		name = getCallerName()
	}
	defer span.Finish()
	if len(args) > 0 {
		span.AddLabels("db.args", fmt.Sprintf("%v", namedValueToLabels(args)))
	}
	labels := []string{labelMethod, "QueryContext", labelQuery, name}
	if conn, ok := w.stmt.(driver.StmtQueryContext); ok {
		ts := time.Now()
		rows, err := conn.QueryContext(nctx, args)
		td := time.Since(ts)
		te := td.Seconds()
		if err != nil {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
			span.SetStatus(tracer.SpanStatusError, err.Error())
		} else {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
		}

		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

		if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
			w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "QueryContext", name, td, err)).Log(ctx, w.opts.LoggerLevel)
		}
		return rows, err
	}

	values, err := namedValueToValue(args)
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()

		span.SetStatus(tracer.SpanStatusError, err.Error())

		if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
			w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "QueryContext", name, 0, err)).Log(ctx, w.opts.LoggerLevel)
		}
		return nil, err
	}
	ts := time.Now()
	rows, err := w.Query(values) // nolint:staticcheck
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		span.SetStatus(tracer.SpanStatusError, err.Error())
	} else {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	}

	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "QueryContext", name, td, err)).Log(ctx, w.opts.LoggerLevel)
	}

	return rows, err
}
