package wrapper

import (
	"context"
	"database/sql/driver"
)

// wrapperStmt defines a wrapper for driver.Stmt
type wrapperStmt struct {
	stmt driver.Stmt
	opts Options
}

// Close implements driver.Stmt Close
func (w *wrapperStmt) Close() error {
	return w.stmt.Close()
}

// NumInput implements driver.Stmt NumInput
func (w *wrapperStmt) NumInput() int {
	return w.stmt.NumInput()
}

// Exec implements driver.Stmt Exec
func (w *wrapperStmt) Exec(args []driver.Value) (driver.Result, error) {
	return w.stmt.Exec(args)
}

// Query implements driver.Stmt Query
func (w *wrapperStmt) Query(args []driver.Value) (driver.Rows, error) {
	return w.stmt.Query(args)
}

// ExecContext implements driver.ExecerContext ExecContext
func (w *wrapperStmt) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	nctx, span := w.opts.Tracer.Start(ctx, "ExecContext")
	defer span.Finish()
	if execerContext, ok := w.stmt.(driver.ExecerContext); ok {
		return execerContext.ExecContext(nctx, query, args)
	}
	values, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}
	return w.Exec(values)
}

// QueryContext implements Driver.QueryerContext QueryContext
func (w *wrapperStmt) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (rows driver.Rows, err error) {
	nctx, span := w.opts.Tracer.Start(ctx, "QueryContext")
	defer span.Finish()
	if queryerContext, ok := w.stmt.(driver.QueryerContext); ok {
		rows, err := queryerContext.QueryContext(nctx, query, args)
		return rows, err
	}
	values, err := namedValueToValue(args)
	if err != nil {
		return nil, err
	}
	return w.Query(values)
}
