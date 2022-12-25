package wrapper

import (
	"context"
	"database/sql/driver"
	"fmt"

	"go.unistack.org/micro/v3/tracer"
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
	if execerContext, ok := w.stmt.(driver.ExecerContext); ok {
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
	res, err := w.Exec(values)
	if err != nil {
		span.AddLabels("error", true)
	}
	return res, err
}

// QueryContext implements Driver.QueryerContext QueryContext
func (w *wrapperStmt) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
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
	if queryerContext, ok := w.stmt.(driver.QueryerContext); ok {
		rows, err := queryerContext.QueryContext(nctx, query, args)
		if err != nil {
			span.AddLabels("error", true)
		}
		return rows, err
	}
	values, err := namedValueToValue(args)
	if err != nil {
		if err != nil {
			span.AddLabels("error", true)
		}
		return nil, err
	}
	rows, err := w.Query(values)
	if err != nil {
		span.AddLabels("error", true)
	}
	return rows, err
}
