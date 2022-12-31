package wrapper

import (
	"context"
	"database/sql/driver"
	"fmt"
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
	span.AddLabels("op", "ExecContext")
	if name := getQueryName(ctx); name != "" {
		span.AddLabels("query", name)
	}
	defer span.Finish()
	if len(args) > 0 {
		span.AddLabels("args", fmt.Sprintf("%v", namedValueToLabels(args)))
	}
	if execerContext, ok := w.stmt.(driver.ExecerContext); ok {
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
	span.AddLabels("op", "QueryContext")
	if name := getQueryName(ctx); name != "" {
		span.AddLabels("query", name)
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
		if err != nil {
			span.AddLabels("error", true)
			span.AddLabels("err", err.Error())
		}
		return nil, err
	}
	rows, err := w.Query(values)
	if err != nil {
		span.AddLabels("error", true)
		span.AddLabels("err", err.Error())
	}
	return rows, err
}
