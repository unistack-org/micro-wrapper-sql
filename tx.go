package wrapper

import (
	"database/sql/driver"

	"go.unistack.org/micro/v3/tracer"
)

// wrapperTx defines a wrapper for driver.Tx
type wrapperTx struct {
	tx   driver.Tx
	span tracer.Span
	opts Options
}

// Commit implements driver.Tx Commit
func (w *wrapperTx) Commit() error {
	if w.span != nil {
		defer w.span.Finish()
	}
	return w.tx.Commit()
}

// Rollback implements driver.Tx Rollback
func (w *wrapperTx) Rollback() error {
	if w.span != nil {
		defer w.span.Finish()
	}
	return w.tx.Rollback()
}
