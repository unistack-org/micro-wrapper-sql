package wrapper

import (
	"context"
	"database/sql/driver"
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
	if connBeginTx, ok := w.conn.(driver.ConnBeginTx); ok {
		tx, err := connBeginTx.BeginTx(nctx, opts)
		if err != nil {
			return nil, err
		}
		return &wrapperTx{tx: tx, opts: w.opts, span: span}, nil
	}
	return w.conn.Begin()
}

// PrepareContext implements driver.ConnPrepareContext PrepareContext
func (w *wrapperConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if connPrepareContext, ok := w.conn.(driver.ConnPrepareContext); ok {
		stmt, err := connPrepareContext.PrepareContext(ctx, query)
		if err != nil {
			return nil, err
		}
		return &wrapperStmt{stmt: stmt, opts: w.opts}, nil
	}
	return w.conn.Prepare(query)
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
	defer span.Finish()
	if execerContext, ok := w.conn.(driver.ExecerContext); ok {
		r, err := execerContext.ExecContext(nctx, query, args)
		return r, err
	}
	values, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}
	return w.Exec(query, values)
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
func (w *wrapperConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (rows driver.Rows, err error) {
	nctx, span := w.opts.Tracer.Start(ctx, "QueryContext")
	defer span.Finish()
	if queryerContext, ok := w.conn.(driver.QueryerContext); ok {
		rows, err := queryerContext.QueryContext(nctx, query, args)
		return rows, err
	}
	values, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}
	return w.Query(query, values)
}
