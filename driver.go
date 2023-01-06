package wrapper

import (
	"context"
	"database/sql/driver"
	"time"
)

// wrapperDriver defines a wrapper for driver.Driver
type wrapperDriver struct {
	driver driver.Driver
	opts   Options
}

// NewWrapper creates and returns a new SQL driver with passed capabilities
func NewWrapper(d driver.Driver, opts ...Option) driver.Driver {
	return &wrapperDriver{driver: d, opts: NewOptions(opts...)}
}

// Open implements driver.Driver Open
func (w *wrapperDriver) Open(name string) (driver.Conn, error) {
	ts := time.Now()
	c, err := w.driver.Open(name)
	td := time.Since(ts)

	if w.opts.LoggerEnabled {
		w.opts.Logger.Fields(w.opts.LoggerObserver(context.TODO(), "Open", labelUnknown, td, err)...).Log(context.TODO(), w.opts.LoggerLevel)
	}

	if err != nil {
		return nil, err
	}

	return &wrapperConn{conn: c, opts: w.opts}, nil
}
