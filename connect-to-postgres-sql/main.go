package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLmode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLmode)
}
func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "556690",
		Database: "test_db",
		SSLmode:  "disable",
	}
	//open the db using the sql driver
	db, err := sql.Open("postgres", cfg.String())
	checkErr(err)

	//Ping to find if the db works
	err = db.Ping()
	checkErr(err)

	defer db.Close()
	fmt.Println("connected!")

	//print out one query row
	var id int
	var name, email string
	row := db.QueryRow(`
  SELECT * 
  FROM customers
  WHERE id=$1`, 1)

	err = row.Scan(&name, &email, &id)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows")
		} else {
			panic(err)
		}
	}

	fmt.Println(email, id)

	//print out multiple rows
	
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
