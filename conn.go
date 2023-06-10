package wrapper

import (
	"context"
	"database/sql/driver"
	"fmt"
	"time"

	"go.unistack.org/micro/v4/tracer"
)

var (
	_ driver.Conn               = (*wrapperConn)(nil)
	_ driver.ConnBeginTx        = (*wrapperConn)(nil)
	_ driver.ConnPrepareContext = (*wrapperConn)(nil)
	_ driver.Pinger             = (*wrapperConn)(nil)
	_ driver.Validator          = (*wrapperConn)(nil)
	_ driver.Queryer            = (*wrapperConn)(nil)
	_ driver.QueryerContext     = (*wrapperConn)(nil)
	_ driver.Execer             = (*wrapperConn)(nil)
	_ driver.ExecerContext      = (*wrapperConn)(nil)
	//	_ driver.Connector
	//	_ driver.Driver
	//	_ driver.DriverContext
)

// wrapperConn defines a wrapper for driver.Conn
type wrapperConn struct {
	d     *wrapperDriver
	dname string
	conn  driver.Conn
	opts  Options
	ctx   context.Context
	span  tracer.Span
}

// Close implements driver.Conn Close
func (w *wrapperConn) Close() error {
	var ctx context.Context
	if w.ctx != nil {
		ctx = w.ctx
	} else {
		ctx = context.Background()
	}

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

	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "Close", labelUnknown, td, err)...).Log(ctx, w.opts.LoggerLevel)
	}

	return err
}

// Begin implements driver.Conn Begin
func (w *wrapperConn) Begin() (driver.Tx, error) {
	var ctx context.Context
	if w.ctx != nil {
		ctx = w.ctx
	} else {
		ctx = context.Background()
	}

	labels := []string{labelMethod, "Begin"}
	ts := time.Now()
	tx, err := w.conn.Begin() // nolint:staticcheck
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
		if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
			w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "Begin", labelUnknown, td, err)...).Log(ctx, w.opts.LoggerLevel)
		}
		return nil, err
	}
	w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "Begin", labelUnknown, td, err)...).Log(ctx, w.opts.LoggerLevel)
	}

	return &wrapperTx{tx: tx, opts: w.opts, ctx: ctx}, nil
}

// BeginTx implements driver.ConnBeginTx BeginTx
func (w *wrapperConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	nctx, span := w.opts.Tracer.Start(ctx, "Transaction")
	span.AddLabels("method", "BeginTx")
	name := getQueryName(ctx)
	if name != "" {
		span.AddLabels("query", name)
	} else {
		name = labelUnknown
	}
	labels := []string{labelMethod, "BeginTx", labelQuery, name}

	connBeginTx, ok := w.conn.(driver.ConnBeginTx)
	if !ok {
		return w.Begin()
	}

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
		if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
			w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "BeginTx", name, td, err)...).Log(ctx, w.opts.LoggerLevel)
		}
		return nil, err
	}
	w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "BeginTx", name, td, err)...).Log(ctx, w.opts.LoggerLevel)
	}
	return &wrapperTx{tx: tx, opts: w.opts, ctx: ctx, span: span}, nil
}

// Prepare implements driver.Conn Prepare
func (w *wrapperConn) Prepare(query string) (driver.Stmt, error) {
	var ctx context.Context
	if w.ctx != nil {
		ctx = w.ctx
	} else {
		ctx = context.Background()
	}

	labels := []string{labelMethod, "Prepare", labelQuery, labelUnknown}
	ts := time.Now()
	stmt, err := w.conn.Prepare(query)
	td := time.Since(ts)
	te := td.Seconds()
	if err != nil {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelFailure)...).Inc()
		w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
		w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

		if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
			w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "Prepare", labelUnknown, td, err)...).Log(ctx, w.opts.LoggerLevel)
		}
		return nil, err
	}
	w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "Prepare", labelUnknown, td, err)...).Log(ctx, w.opts.LoggerLevel)
	}

	return wrapStmt(stmt, query, w.opts), nil
}

// PrepareContext implements driver.ConnPrepareContext PrepareContext
func (w *wrapperConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	var nctx context.Context
	var span tracer.Span
	if w.ctx != nil {
		nctx, span = w.opts.Tracer.Start(w.ctx, "PrepareContext")
	} else {
		nctx, span = w.opts.Tracer.Start(ctx, "PrepareContext")
	}
	span.AddLabels("method", "PrepareContext")
	name := getQueryName(ctx)
	if name != "" {
		span.AddLabels("query", name)
	} else {
		name = labelUnknown
	}

	labels := []string{labelMethod, "PrepareContext", labelQuery, name}
	conn, ok := w.conn.(driver.ConnPrepareContext)
	if !ok {
		return w.Prepare(query)
	}

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
		if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
			w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "PrepareContext", name, td, err)...).Log(ctx, w.opts.LoggerLevel)
		}
		return nil, err
	}
	w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)
	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "PrepareContext", name, td, err)...).Log(ctx, w.opts.LoggerLevel)
	}

	return wrapStmt(stmt, query, w.opts), nil
}

// Exec implements driver.Execer Exec
func (w *wrapperConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	var ctx context.Context
	if w.ctx != nil {
		ctx = w.ctx
	} else {
		ctx = context.Background()
	}
	labels := []string{labelMethod, "Exec", labelQuery, labelUnknown}

	// nolint:staticcheck
	conn, ok := w.conn.(driver.Execer)
	if !ok {
		return nil, driver.ErrSkip
	}

	ts := time.Now()
	res, err := conn.Exec(query, args)
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
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "Exec", labelUnknown, td, err)...).Log(ctx, w.opts.LoggerLevel)
	}
	return res, err
}

// Exec implements driver.StmtExecContext ExecContext
func (w *wrapperConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	var nctx context.Context
	var span tracer.Span

	if w.ctx != nil {
		nctx, span = w.opts.Tracer.Start(w.ctx, "ExecContext")
	} else {
		nctx, span = w.opts.Tracer.Start(ctx, "ExecContext")
	}
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
	fmt.Printf("EXECCONTETX args %#+v\n", args)

	conn, ok := w.conn.(driver.ExecerContext)
	if !ok {
		// nolint:staticcheck
		return nil, driver.ErrSkip
	}

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
	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "ExecContext", name, td, err)...).Log(ctx, w.opts.LoggerLevel)
	}

	return res, err
}

// Ping implements driver.Pinger Ping
func (w *wrapperConn) Ping(ctx context.Context) error {
	conn, ok := w.conn.(driver.Pinger)

	if !ok {
		// fallback path to check db alive
		pc, err := w.d.Open(w.dname)
		if err != nil {
			return err
		}
		return pc.Close()
	}

	var nctx context.Context
	var span tracer.Span
	if w.ctx != nil {
		nctx, span = w.opts.Tracer.Start(w.ctx, "Ping")
	} else {
		nctx, span = w.opts.Tracer.Start(ctx, "Ping")
	}
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
		if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
			w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "Ping", labelUnknown, td, err)...).Log(ctx, w.opts.LoggerLevel)
		}
		return err
	} else {
		w.opts.Meter.Counter(meterRequestTotal, append(labels, labelStatus, labelSuccess)...).Inc()
	}
	w.opts.Meter.Summary(meterRequestLatencyMicroseconds, labels...).Update(te)
	w.opts.Meter.Histogram(meterRequestDurationSeconds, labels...).Update(te)

	return nil
}

// Query implements driver.Queryer Query
func (w *wrapperConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	var ctx context.Context
	if w.ctx != nil {
		ctx = w.ctx
	} else {
		ctx = context.Background()
	}
	// nolint:staticcheck
	conn, ok := w.conn.(driver.Queryer)
	if !ok {
		return nil, driver.ErrSkip
	}

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
	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "Query", labelUnknown, td, err)...).Log(ctx, w.opts.LoggerLevel)
	}
	return rows, err
}

// QueryContext implements Driver.QueryerContext QueryContext
func (w *wrapperConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	var nctx context.Context
	var span tracer.Span
	if w.ctx != nil {
		nctx, span = w.opts.Tracer.Start(w.ctx, "QueryContext")
	} else {
		nctx, span = w.opts.Tracer.Start(ctx, "QueryContext")
	}
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
	conn, ok := w.conn.(driver.QueryerContext)
	if !ok {
		return nil, driver.ErrSkip
	}

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
	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(ctx, "QueryContext", name, td, err)...).Log(ctx, w.opts.LoggerLevel)
	}

	return rows, err
}

// CheckNamedValue implements driver.NamedValueChecker
func (w *wrapperConn) CheckNamedValue(v *driver.NamedValue) error {
	s, ok := w.conn.(driver.NamedValueChecker)
	if !ok {
		return driver.ErrSkip
	}
	return s.CheckNamedValue(v)
}

// IsValid implements driver.Validator
func (w *wrapperConn) IsValid() bool {
	v, ok := w.conn.(driver.Validator)
	if !ok {
		return w.conn != nil
	}
	return v.IsValid()
}

func (w *wrapperConn) ResetSession(ctx context.Context) error {
	s, ok := w.conn.(driver.SessionResetter)
	if !ok {
		return driver.ErrSkip
	}
	return s.ResetSession(ctx)
}
