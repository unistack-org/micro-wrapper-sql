package wrapper

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

// ErrUnsupported is an error returned when the underlying driver doesn't provide a given function.
var ErrUnsupported = errors.New("operation unsupported by the underlying driver")

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

// namedValueToLabels convert driver arguments to interface{} slice
func namedValueToLabels(named []driver.NamedValue) []interface{} {
	largs := make([]interface{}, len(named)*2)
	var name string
	for _, param := range named {
		if param.Name != "" {
			name = param.Name
		} else {
			name = fmt.Sprintf("$%d", param.Ordinal)
		}

		largs = append(largs, name, param.Value)
	}
	return largs
}
