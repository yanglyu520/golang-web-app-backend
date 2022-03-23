package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"net/http"
)

var db *sql.DB
var err error

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLmode  string
}

type User struct {
	Id    int
	Name  string
	Email string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLmode)
}

func init() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "556690",
		Database: "test_db",
		SSLmode:  "disable",
	}

	db, err = sql.Open("postgres", cfg.String())
	checkErr(err)

	err = db.Ping()
	checkErr(err)

	fmt.Println("db is connected!\n")
}

func main() {
	defer db.Close()

  http.Handle("/", http.HandlerFunc(index))
	http.Handle("/s", http.HandlerFunc(singleRow))
	http.Handle("/m", http.HandlerFunc(multipleRows))

	http.ListenAndServe(":8080", nil)

}

func singleRow(w http.ResponseWriter, r *http.Request) {
	user := User{}
	row := db.QueryRow(`SELECT * FROM customers WHERE id=$1`, 1)

	err = row.Scan(&user.Name, &user.Email, &user.Id)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("no rows")
		} else {
			panic(err)
		}
	}
	fmt.Fprintf(w, "%d, %s, %s", user.Id, user.Name, user.Email)

}

func multipleRows(w http.ResponseWriter, r *http.Request) {
	user := User{}
	var users []User

	printRequest(r)

	rows, err := db.Query(`SELECT * FROM customers`)
	checkErr(err)

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.Name, &user.Email, &user.Id)
		checkErr(err)
		users = append(users, user)
	}

	for _, v := range users {
		fmt.Fprintf(w, "%d, %s, %s", v.Id, v.Name, v.Email)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	printRequest(r)
	_, err := io.WriteString(w, "at index")
	checkErr(err)
}

func printRequest(r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
}
