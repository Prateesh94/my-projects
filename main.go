package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "postgres"
)

func main() {
	sqlcon := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Println(sqlcon)
	db, er := sql.Open("postgres", sqlcon)
	db.Ping()
	fmt.Println(er, " is established")
	qry := `insert into users(name,email,password) values($1,$2,$3)`
	db.Exec(qry, "you are", "myfire", "the one desire")
	fmt.Println("Added User")
}
