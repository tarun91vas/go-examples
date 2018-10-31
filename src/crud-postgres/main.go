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
	connStr := "postgres://user:pass@localhost/postgres?sslmode=disable"
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

	// db.QueryRow example: select one row
	row := db.QueryRow("SELECT id, username, email FROM pgUser WHERE username = $1", "test")
	err = row.Scan(&userRow.id, &userRow.name, &userRow.email)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("User not found")
	case err != nil:
		fmt.Println(err)
	}
	fmt.Println(userRow)

	// db.Exec example: Create, update and delete operations
	nr, err := db.Exec("UPDATE pgUser SET phone = $1 WHERE username = $2", "9843465431", "test-1")
	if err != nil {
		fmt.Println("update failed: ", err)
	}
	numRows, _ := nr.RowsAffected()
	fmt.Println("Rows affected: ", numRows)

}
