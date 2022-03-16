# How to connect golang web app to postgres sql

## 1. install postgres and pgadmin on mac
1. Follow this instruction [here](https://www.enterprisedb.com/postgres-tutorials/installation-postgresql-mac-os)

2. You can install the cli to postgres sql using homebrew and then run this
- `sudo -u postgres psql`
- `\l` to list all the db
- `\c <db name>` to connect to a db
- `SELECT current_database();` to find the current db
- `SELECT current_user;` to find the current user
- `DROP DATABASE xx;` to delete a db
- `\q` to quit

3. Here are some useful SQL scripts:

```sql
CREATE DATABASE xxx
...
```

## 2. create a table with some demo data
use the pgadmin or SQL script to create a test table and some columns and some test data

## 3. Functions to use on sql package
Functions to test connections to a DB:
- `func Open(driverName, dataSourceName string) (*DB, error)`
- `func (db *DB) Ping() error`

Functions to run query to a db:
- functions does not return data, use `Exec`

- `QueryRow`, `func (db *DB) QueryRow(query string, args ...any) *Row` returns at most single row
- `Query`, `func (db *DB) Query(query string, args ...any) (*Rows, error)` will return multiple rows, using `Rows.Next`, `func (rs \*Rows) Next() bool` to tell if it reaches the end of the rows
- `func (r \*Row) Scan(dest ...any) error` 


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
var id int
var name, email string
row := db.QueryRow(`
  SELECT * 
  FROM customers
  WHERE id=$1`,1)

err = row.Scan(&id, &name, &email)

if err != nil {
  if err == sql.ErrNoRows {
    fmt.Println("no rows")
  }else {
    panic(err)
  }
}

//query for multiple rows
```
---
#### Reference
- [go doc on sql package](https://pkg.go.dev/database/sql), Note: The sql package must be used in conjunction with a database driver. See https://golang.org/s/sqldrivers for a list of drivers.
- [example provided in go doc](https://github.com/golang/go/wiki/SQLInterface)
- functions return data, need to follow this [instruction](https://go.dev/doc/database/querying)

#### What is left to learn
1. More on connect to postgres sql with cli
2. Maybe use docker to connect to posgres sql local, or aws postgres sql remotely