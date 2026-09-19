package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "app.db")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = db.Close()
	}()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS people (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			age INTEGER NOT NULL
	)`)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table ready")

	res, err := db.Exec("INSERT INTO people (name, age) VALUES (?, ?)", "Alice", 30)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("inserted id: ", id)

	rows, err := db.Query("SELECT id, name, age FROM people")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var pid int
		var name string
		var age int

		if err := rows.Scan(&pid, &name, &age); err != nil {
			log.Fatal(err)
		}

		fmt.Println(pid, name, age)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	var found string
	err = db.QueryRow("SELECT name FROM people WHERE id = ?", 999).Scan(&found)

	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no person with id 999")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("found:", found)
	}
}
