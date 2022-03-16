# How to connect golang web app to postgres sql

## install postgres and pgadmin on mac
Follow this instruction here
https://www.enterprisedb.com/postgres-tutorials/installation-postgresql-mac-os

## create a table with pgadmin
use the pgadmin tool to create a test table and some columns and some test data


## go doc 
- [go doc](https://pkg.go.dev/database/sql)
- The sql package must be used in conjunction with a database driver. See https://golang.org/s/sqldrivers for a list of drivers.
- [example](https://github.com/golang/go/wiki/SQLInterface)

functions available are
- func Open(driverName, dataSourceName string) (*DB, error)
- func (db *DB) Ping() error


```go
//connecting to a db
db, err := sql.Open(driver, dataSourceName)
//check if it is working using ping
if err := db.PingContext(ctx); err != nil {
log.Fatal(err)
}
if err := db.Ping(); err != nil {
log.Fatal(err)
}

//query for a single row

//query for multiple rows
```

---
#### What is left to learn
1. More on connect to postgres sql with cli
2. Maybe use docker to connect to posgres sql, or aws postgres sql