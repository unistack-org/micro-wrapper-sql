# micro-wrapper-sql

Example for For postgres 

```go
package postgres

import (
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func Connect(cfg *PostgresConfig) (*sqlx.DB, error) {
	// format connection string
	cstr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable&statement_cache_mode=describe",
		cfg.Login,
		url.QueryEscape(cfg.Passw),
		cfg.Addr,
		cfg.DBName,
	)

	// parse connection string
	dbConf, err := pgx.ParseConfig(cstr)
	if err != nil {
		return nil, err
	}

	// needed for pgbouncer
	dbConf.RuntimeParams = map[string]string{
		"standard_conforming_strings": "on",
		"application_name":            cfg.AppName,
	}
	// may be needed for pbbouncer, needs to check
	//dbConf.PreferSimpleProtocol = true
	// register pgx conn
	dsn := stdlib.RegisterConnConfig(dbConf)


  sql.Register("micro-wrapper-sql", wrapper.WrapDriver(
    &stdlib.Driver{}, 
    wrapper.Tracer(some.NewTracer()),
  ))
  wdb, err := sql.Open("micro-wrapper-sql", dsn)
	if err != nil {
		return nil, err
	}

	db := sqlx.NewDb(wdb, "pgx")
	db.SetMaxOpenConns(int(cfg.ConnMax))
	db.SetMaxIdleConns(int(cfg.ConnMaxIdle))
	db.SetConnMaxLifetime(time.Duration(cfg.ConnLifetime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)
	
	return db, nil
}
```
