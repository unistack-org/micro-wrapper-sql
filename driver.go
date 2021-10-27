package wrapper

import (
	"database/sql/driver"
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
	c, err := w.driver.Open(name)
	if err != nil {
		return nil, err
	}
	return &wrapperConn{conn: c, opts: w.opts}, nil
}
