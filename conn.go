package wrapper

import (
	"context"
	"database/sql/driver"
	"fmt"
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
	nctx, span := w.opts.Tracer.Start(ctx, "BeginTx")
	span.AddLabels("op", "BeginTx")
	if name := getQueryName(ctx); name != "" {
		span.AddLabels("query", name)
	}
	if connBeginTx, ok := w.conn.(driver.ConnBeginTx); ok {
		tx, err := connBeginTx.BeginTx(nctx, opts)
		if err != nil {
			span.AddLabels("error", true)
			span.AddLabels("err", err.Error())
			return nil, err
		}
		return &wrapperTx{tx: tx, opts: w.opts, span: span}, nil
	}
	tx, err := w.conn.Begin()
	if err != nil {
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
	}
	return tx, err
}

// PrepareContext implements driver.ConnPrepareContext PrepareContext
func (w *wrapperConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	nctx, span := w.opts.Tracer.Start(ctx, "PrepareContext")
	span.AddLabels("op", "PrepareContext")
	if name := getQueryName(ctx); name != "" {
		span.AddLabels("query", name)
	}
	if connPrepareContext, ok := w.conn.(driver.ConnPrepareContext); ok {
		stmt, err := connPrepareContext.PrepareContext(nctx, query)
		if err != nil {
			span.AddLabels("error", true)
			span.AddLabels("err", err.Error())
			return nil, err
		}
		return &wrapperStmt{stmt: stmt, opts: w.opts}, nil
	}
	stmt, err := w.conn.Prepare(query)
	if err != nil {
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
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
	nctx, span := w.opts.Tracer.Start(ctx, "ExecContext")
	span.AddLabels("op", "ExecContext")
	if name := getQueryName(ctx); name != "" {
		span.AddLabels("query", name)
	}
	defer span.Finish()
	if len(args) > 0 {
		span.AddLabels("args", fmt.Sprintf("%v", namedValueToLabels(args)))
	}
	if execerContext, ok := w.conn.(driver.ExecerContext); ok {
		res, err := execerContext.ExecContext(nctx, query, args)
		if err != nil {
			span.AddLabels("error", true)
			span.AddLabels("err", err.Error())
		}
		return res, err
	}
	values, err := namedValueToValue(args)
	if err != nil {
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
		return nil, err
	}
	res, err := w.Exec(query, values)
	if err != nil {
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
	}
	return res, err
}

// Ping implements driver.Pinger Ping
func (w *wrapperConn) Ping(ctx context.Context) error {
	if pinger, ok := w.conn.(driver.Pinger); ok {
		nctx, span := w.opts.Tracer.Start(ctx, "Ping")
		defer span.Finish()
		err := pinger.Ping(nctx)
		if err != nil {
			span.AddLabels("error", true)
			span.AddLabels("err", err.Error())
			return err
		}
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
	nctx, span := w.opts.Tracer.Start(ctx, "QueryContext")
	span.AddLabels("op", "QueryContext")
	if name := getQueryName(ctx); name != "" {
		span.AddLabels("query", name)
	}
	defer span.Finish()
	if len(args) > 0 {
		span.AddLabels("args", fmt.Sprintf("%v", namedValueToLabels(args)))
	}
	if queryerContext, ok := w.conn.(driver.QueryerContext); ok {
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
	rows, err := w.Query(query, values)
	if err != nil {
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
	}
	return rows, err
}
