package wrapper

import (
	"context"
	"database/sql/driver"
	"time"

	"go.unistack.org/micro/v3/tracer"
)

var _ driver.Tx = (*wrapperTx)(nil)

// wrapperTx defines a wrapper for driver.Tx
type wrapperTx struct {
	tx   driver.Tx
	span tracer.Span
	opts Options
	ctx  context.Context
}

// Commit implements driver.Tx Commit
func (w *wrapperTx) Commit() error {
	ts := time.Now()
	err := w.tx.Commit()
	td := time.Since(ts)

	if w.span != nil {
		if err != nil {
			w.span.AddLabels("error", true)
			w.span.AddLabels("err", err.Error())
		}
		w.span.Finish()
	}

	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(w.ctx, "Commit", labelUnknown, td, err)...).Log(w.ctx, w.opts.LoggerLevel)
	}

	w.ctx = nil

	return err
}

// Rollback implements driver.Tx Rollback
func (w *wrapperTx) Rollback() error {
	ts := time.Now()
	err := w.tx.Rollback()
	td := time.Since(ts)

	if w.span != nil {
		if err != nil {
			w.span.AddLabels("error", true)
			w.span.AddLabels("err", err.Error())
		}
		w.span.Finish()
	}

	if w.opts.LoggerEnabled && w.opts.Logger.V(w.opts.LoggerLevel) {
		w.opts.Logger.Fields(w.opts.LoggerObserver(w.ctx, "Rollback", labelUnknown, td, err)...).Log(w.ctx, w.opts.LoggerLevel)
	}

	w.ctx = nil

	return err
}
