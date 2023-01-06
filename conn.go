package wrapper

import (
	"context"
	"database/sql/driver"
	"fmt"
	"time"
)

// wrapperConn defines a wrapper for driver.Conn
type wrapperConn struct {
	conn driver.Conn
	opts Options
}

// Prepare implements driver.Conn Prepare
func (w *wrapperConn) Prepare(query string) (driver.Stmt, error) {
	labels := []string{labelMethod, "Prepare", labelQuery, labelUnknown}
	ts := time.Now()
	stmt, err := w.conn.Prepare(query)
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

		if w.opts.LoggerEnabled {
			w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "Prepare", labelUnknown, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
		}
		return nil, err
	}
	w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "QueryContext", labelUnknown, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
	}

	return &wrapperStmt{stmt: stmt, opts: w.opts}, nil
}

// Close implements driver.Conn Close
func (w *wrapperConn) Close() error {
	labels := []string{labelMethod, "Close"}
	ts := time.Now()
	err := w.conn.Close()
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
	} else {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	}
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "QueryContext", labelUnknown, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
	}

	return err
}

// Begin implements driver.Conn Begin
func (w *wrapperConn) Begin() (driver.Tx, error) {
	labels := []string{labelMethod, "Begin"}
	ts := time.Now()
	// nolint:staticcheck
	tx, err := w.conn.Begin()
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
		if w.opts.LoggerEnabled {
			w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "Begin", labelUnknown, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
		}
		return nil, err
	}
	w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "Begin", labelUnknown, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
	}
	return &wrapperTx{tx: tx, opts: w.opts}, nil
}

// BeginTx implements driver.ConnBeginTx BeginTx
func (w *wrapperConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	nctx, span := w.opts.Tracer.Start(ctx, "BeginTx")
	span.AddLabels("method", "BeginTx")
	name := getQueryName(ctx)
	if name != "" {
		span.AddLabels("query", name)
	} else {
		name = labelUnknown
	}
	labels := []string{labelMethod, "BeginTx", labelQuery, name}
	if connBeginTx, ok := w.conn.(driver.ConnBeginTx); ok {
		ts := time.Now()
		tx, err := connBeginTx.BeginTx(nctx, opts)
		td := time.Since(ts)
		te := td.Seconds()
		if err != nil {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
			w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
			w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
			span.AddLabels("error", true)
			span.AddLabels("err", err.Error())
			if w.opts.LoggerEnabled {
				w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "BeginTx", name, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
			}
			return nil, err
		}
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
		if w.opts.LoggerEnabled {
			w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "BeginTx", name, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
		}
		return &wrapperTx{tx: tx, opts: w.opts, span: span}, nil
	}
	ts := time.Now()
	// nolint:staticcheck
	tx, err := w.conn.Begin()
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
	}
	w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "BeginTx", name, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
	}
	return tx, nil
}

// PrepareContext implements driver.ConnPrepareContext PrepareContext
func (w *wrapperConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	nctx, span := w.opts.Tracer.Start(ctx, "PrepareContext")
	span.AddLabels("method", "PrepareContext")
	name := getQueryName(ctx)
	if name != "" {
		span.AddLabels("query", name)
	} else {
		name = labelUnknown
	}
	labels := []string{labelMethod, "PrepareContext", labelQuery, name}
	if conn, ok := w.conn.(driver.ConnPrepareContext); ok {
		ts := time.Now()
		stmt, err := conn.PrepareContext(nctx, query)
		td := time.Since(ts)
		te := td.Seconds()
		if err != nil {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
			w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
			w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
			span.AddLabels("error", true)
			span.AddLabels("err", err.Error())
			if w.opts.LoggerEnabled {
				w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "PrepareContext", name, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
			}
			return nil, err
		}
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
		if w.opts.LoggerEnabled {
			w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "PrepareContext", name, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
		}
		return &wrapperStmt{stmt: stmt, opts: w.opts}, nil
	}
	ts := time.Now()
	stmt, err := w.conn.Prepare(query)
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
	}
	w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "PrepareContext", name, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
	}
	return stmt, nil
}

// Exec implements driver.Execer Exec
func (w *wrapperConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	// nolint:staticcheck
	labels := []string{labelMethod, "Exec", labelQuery, labelUnknown}
	if execer, ok := w.conn.(driver.Execer); ok {
		ts := time.Now()
		res, err := execer.Exec(query, args)
		td := time.Since(ts)
		te := td.Seconds()
		if err != nil {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		} else {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
		}
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
		if w.opts.LoggerEnabled {
			w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "Exec", labelUnknown, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
		}
		return res, err
	}
	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "Exec", labelUnknown, 0, ErrUnsupported)...).Log(context.TODO(), w.opts.LoggerLevel)
	}
	return nil, ErrUnsupported
}

// Exec implements driver.StmtExecContext ExecContext
func (w *wrapperConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
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
	if conn, ok := w.conn.(driver.ExecerContext); ok {
		ts := time.Now()
		res, err := conn.ExecContext(nctx, query, args)
		td := time.Since(ts)
		te := td.Seconds()
		if err != nil {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
			span.AddLabels("error", true)
			span.AddLabels("err", err.Error())
		} else {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
		}
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
		if w.opts.LoggerEnabled {
			w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "ExecContext", name, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
		}
		return res, err
	}
	values, err := namedValueToValue(args)
	if err != nil {
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
		if w.opts.LoggerEnabled {
			w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "ExecContext", labelUnknown, 0, err)...).Log(context.TODO(), w.opts.LoggerLevel)
		}
		return nil, err
	}
	ts := time.Now()
	// nolint:staticcheck
	res, err := w.Exec(query, values)
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
	} else {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	}
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "ExecContext", name, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
	}
	return res, err
}

// Ping implements driver.Pinger Ping
func (w *wrapperConn) Ping(ctx context.Context) error {
	if conn, ok := w.conn.(driver.Pinger); ok {
		nctx, span := w.opts.Tracer.Start(ctx, "Ping")
		defer span.Finish()
		labels := []string{labelMethod, "Ping"}
		ts := time.Now()
		err := conn.Ping(nctx)
		td := time.Since(ts)
		te := td.Seconds()
		if err != nil {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
			span.AddLabels("error", true)
			span.AddLabels("err", err.Error())
			if w.opts.LoggerEnabled {
				w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "Ping", labelUnknown, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
			}
			return err
		} else {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
		}
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
	}
	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "Ping", labelUnknown, 0, ErrUnsupported)...).Log(context.TODO(), w.opts.LoggerLevel)
	}
	return ErrUnsupported
}

// Query implements driver.Queryer Query
func (w *wrapperConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	// nolint:staticcheck
	if conn, ok := w.conn.(driver.Queryer); ok {
		labels := []string{labelMethod, "Query", labelQuery, labelUnknown}
		ts := time.Now()
		rows, err := conn.Query(query, args)
		td := time.Since(ts)
		te := td.Seconds()
		if err != nil {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		} else {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
		}
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
		if w.opts.LoggerEnabled {
			w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "Query", labelUnknown, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
		}
		return rows, err
	}
	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "Query", labelUnknown, 0, ErrUnsupported)...).Log(context.TODO(), w.opts.LoggerLevel)
	}
	return nil, ErrUnsupported
}

// QueryContext implements Driver.QueryerContext QueryContext
func (w *wrapperConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
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
	labels := []string{labelMethod, "QueryContext", labelQuery, name}
	if conn, ok := w.conn.(driver.QueryerContext); ok {
		ts := time.Now()
		rows, err := conn.QueryContext(nctx, query, args)
		td := time.Since(ts)
		te := td.Seconds()
		if err != nil {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
			span.AddLabels("error", true)
			span.AddLabels("err", err.Error())
		} else {
			w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
		}
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
		if w.opts.LoggerEnabled {
			w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "QueryContext", name, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
		}
		return rows, err
	}
	values, err := namedValueToValue(args)
	if err != nil {
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
		if w.opts.LoggerEnabled {
			w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "QueryContext", name, 0, err)...).Log(context.TODO(), w.opts.LoggerLevel)
		}
		return nil, err
	}
	ts := time.Now()
	// nolint:staticcheck
	rows, err := w.Query(query, values)
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
	} else {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	}
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "QueryContext", name, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
	}
	return rows, err
}
