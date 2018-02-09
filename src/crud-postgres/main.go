package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type user struct {
	id    int
	name  string
	email string
}

func main() {
	// Get connection
	//todo insert user pass
	connStr := "postgres://:@localhost/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	// db.Query exmaple: select all
	var userRow user
	rows, err := db.Query("SELECT id, username, email FROM pgUser")
	if err != nil {
		fmt.Println("Query failed: " + err.Error())
	}
	for rows.Next() {
		rows.Scan(&userRow.id, &userRow.name, &userRow.email)
		fmt.Println(userRow)
	}

	//err := row.Scan()

	// username := "test"
	// rows, err := db.Query("SELECT id, username, email FROM pgUser WHERE username = $1", username)
	// if err != nil {
	// 	fmt.Println("Query failed: " + err.Error())
	// }

	// var userRow user
	// for rows.Next() {
	// 	rows.Scan(&userRow.id, &userRow.username, &userRow.email)
	// 	fmt.Println(userRow)
	// }

}
