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

Functions to run query and return data:
1. single row
- `QueryRow`, `func (db *DB) QueryRow(query string, args ...any) *Row` return pointer to a single row
- then call `Scan`, `func (r \*Row) Scan(dest ...any) error` will take a row and scan each field into golang var

2. multiple rows
- `Query`, `func (db *DB) Query(query string, args ...any) (*Rows, error)` will return the pointer to the rows(perhaps a linked list or tree)
- use `Rows.Next`, `func (rs \*Rows) Next() bool` for iteration
- need `rows.Close()` to close 
- `rows.Err() error` returns the error, if any, that was encountered during iteration. Err may be called after an explicit or implicit Close.



Functions does not return data
- `Exec`, `

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

row := db.QueryRow(`SELECT * FROM customers WHERE id=$1`, 1)
err = row.Scan(&name, &email, &id)

if err != nil {
  if err == sql.ErrNoRows {
    fmt.Println("no rows")
  } else {
    panic(err)
  }
}

fmt.Println(email, id)

//query for multiple rows
//print out multiple rows
rows, err := db.Query(`SELECT * FROM customers`)
checkErr(err)
defer rows.Close()
//var to store data
s := "RETRIEVED RECORDS: \n"

for rows.Next() {
  err = rows.Scan(&name, &email, &id)
  checkErr(err)
  s += name + " " + email + "\n"
}

fmt.Println(s)


```
---
#### Reference
- [go doc on sql package](https://pkg.go.dev/database/sql), Note: The sql package must be used in conjunction with a database driver. See https://golang.org/s/sqldrivers for a list of drivers.
- [example provided in go doc](https://github.com/golang/go/wiki/SQLInterface)
- functions return data, need to follow this [instruction](https://go.dev/doc/database/querying)

#### What is left to learn
1. More on connect to postgres sql with cli
2. Maybe use docker to connect to posgres sql local, or aws postgres sql remotely