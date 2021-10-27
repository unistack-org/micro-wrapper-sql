package wrapper

import (
	"database/sql/driver"
	"errors"
)

// ErrUnsupported is an error returned when the underlying driver doesn't provide a given function.
var ErrUnsupported = errors.New("operation unsupported by the underlying driver")

/*
// newSpan creates a new opentracing.Span instance from the given context.
func (t *tracer) newSpan(ctx context.Context) opentracing.Span {
	name := t.nameFunc(ctx)
	var opts []opentracing.StartSpanOption
	parent := opentracing.SpanFromContext(ctx)
	if parent != nil {
		opts = append(opts, opentracing.ChildOf(parent.Context()))
	}
	span := t.t.StartSpan(name, opts...)
	return span
}
*/

// namedValueToValue converts driver arguments of NamedValue format to Value format. Implemented in the same way as in
// database/sql ctxutil.go.
func namedValueToValue(named []driver.NamedValue) ([]driver.Value, error) {
	dargs := make([]driver.Value, len(named))
	for n, param := range named {
		if len(param.Name) > 0 {
			return nil, errors.New("sql: driver does not support the use of Named Parameters")
		}
		dargs[n] = param.Value
	}
	return dargs, nil
}
