package wrapper

import (
	"context"
	"database/sql/driver"
	"fmt"

	"go.unistack.org/micro/v3/tracer"
)

// wrapperConn defines a wrapper for driver.Conn
type wrapperConn struct {
	conn driver.Conn
	opts Options
}

// Prepare implements driver.Conn Prepare
func (w *wrapperConn) Prepare(query string) (driver.Stmt, error) {
	stmt, err := w.conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	return &wrapperStmt{stmt: stmt, opts: w.opts}, nil
}

// Close implements driver.Conn Close
func (w *wrapperConn) Close() error {
	return w.conn.Close()
}

// Begin implements driver.Conn Begin
func (w *wrapperConn) Begin() (driver.Tx, error) {
	tx, err := w.conn.Begin()
	if err != nil {
		return nil, err
	}
	return &wrapperTx{tx: tx, opts: w.opts}, nil
}

// BeginTx implements driver.ConnBeginTx BeginTx
func (w *wrapperConn) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	var nctx context.Context
	var span tracer.Span

	name := getQueryName(ctx)
	if name != "" {
		nctx, span = w.opts.Tracer.Start(ctx, "BeginTx "+name)
	} else {
		nctx, span = w.opts.Tracer.Start(ctx, "BeginTx")
	}
	if name == "" {
		name = "unknown"
	}
	span.AddLabels("query", name)
	if connBeginTx, ok := w.conn.(driver.ConnBeginTx); ok {
		tx, err := connBeginTx.BeginTx(nctx, opts)
		if err != nil {
			span.AddLabels("error", true)
			return nil, err
		}
		return &wrapperTx{tx: tx, opts: w.opts, span: span}, nil
	}
	tx, err := w.conn.Begin()
	if err != nil {
		span.AddLabels("error", true)
	}
	return tx, err
}

// PrepareContext implements driver.ConnPrepareContext PrepareContext
func (w *wrapperConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	var nctx context.Context
	var span tracer.Span

	name := getQueryName(ctx)
	if name != "" {
		nctx, span = w.opts.Tracer.Start(ctx, "BeginTx "+name)
	} else {
		nctx, span = w.opts.Tracer.Start(ctx, "BeginTx")
	}
	if name == "" {
		name = "unknown"
	}
	span.AddLabels("query", name)
	if connPrepareContext, ok := w.conn.(driver.ConnPrepareContext); ok {
		stmt, err := connPrepareContext.PrepareContext(nctx, query)
		if err != nil {
			span.AddLabels("error", true)
			return nil, err
		}
		return &wrapperStmt{stmt: stmt, opts: w.opts}, nil
	}
	stmt, err := w.conn.Prepare(query)
	if err != nil {
		span.AddLabels("error", true)
	}
	return stmt, err
}

// Exec implements driver.Execer Exec
func (w *wrapperConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	if execer, ok := w.conn.(driver.Execer); ok {
		return execer.Exec(query, args)
	}
	return nil, ErrUnsupported
}

// Exec implements driver.StmtExecContext ExecContext
func (w *wrapperConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	var nctx context.Context
	var span tracer.Span
	name := getQueryName(ctx)
	if name != "" {
		nctx, span = w.opts.Tracer.Start(ctx, "ExecContext "+name)
	} else {
		nctx, span = w.opts.Tracer.Start(ctx, "ExecContext")
	}
	defer span.Finish()
	if name == "" {
		name = "unknown"
	}
	span.AddLabels("args", fmt.Sprintf("%v", namedValueToLabels(args)))
	span.AddLabels("query", name)
	if execerContext, ok := w.conn.(driver.ExecerContext); ok {
		res, err := execerContext.ExecContext(nctx, query, args)
		if err != nil {
			span.AddLabels("error", true)
		}
		return res, err
	}
	values, err := namedValueToValue(args)
	if err != nil {
		span.AddLabels("error", true)
		return nil, err
	}
	res, err := w.Exec(query, values)
	if err != nil {
		span.AddLabels("error", true)
	}
	return res, err
}

// Ping implements driver.Pinger Ping
func (w *wrapperConn) Ping(ctx context.Context) error {
	if pinger, ok := w.conn.(driver.Pinger); ok {
		nctx, span := w.opts.Tracer.Start(ctx, "Ping")
		defer span.Finish()
		return pinger.Ping(nctx)
	}
	return ErrUnsupported
}

// Query implements driver.Queryer Query
func (w *wrapperConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	if queryer, ok := w.conn.(driver.Queryer); ok {
		return queryer.Query(query, args)
	}
	return nil, ErrUnsupported
}

// QueryContext implements Driver.QueryerContext QueryContext
func (w *wrapperConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	var nctx context.Context
	var span tracer.Span
	name := getQueryName(ctx)
	if name != "" {
		nctx, span = w.opts.Tracer.Start(ctx, "QueryContext "+name)
	} else {
		nctx, span = w.opts.Tracer.Start(ctx, "QueryContext")
	}
	defer span.Finish()
	if name == "" {
		name = "unknown"
	}
	span.AddLabels("args", fmt.Sprintf("%v", namedValueToLabels(args)))
	span.AddLabels("query", name)
	if queryerContext, ok := w.conn.(driver.QueryerContext); ok {
		rows, err := queryerContext.QueryContext(nctx, query, args)
		if err != nil {
			span.AddLabels("error", true)
		}
		return rows, err
	}
	values, err := namedValueToValue(args)
	if err != nil {
		span.AddLabels("error", true)
		return nil, err
	}
	rows, err := w.Query(query, values)
	if err != nil {
		span.AddLabels("error", true)
	}
	return rows, err
}
