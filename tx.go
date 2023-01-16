package wrapper

import (
	"context"
	"database/sql/driver"
	"time"

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
	ts := time.Now()
	err := w.tx.Commit()
	td := time.Since(ts)

	if err != nil {
		w.span.AddLabels("error", true)
		w.span.AddLabels("err", err.Error())
	}

	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "Commit", labelUnknown, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
	}

	return err
}

// Rollback implements driver.Tx Rollback
func (w *wrapperTx) Rollback() error {
	if w.span != nil {
		defer w.span.Finish()
	}
	ts := time.Now()
	err := w.tx.Rollback()
	td := time.Since(ts)

	if err != nil {
		w.span.AddLabels("error", true)
		w.span.AddLabels("err", err.Error())
	}

	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "Rollback", labelUnknown, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
	}

	return err
}
