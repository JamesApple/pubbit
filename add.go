package main

import (
	"database/sql"
	"log"
)

func add(config Config) {
	db, err := sql.Open("postgres", config.PostgresURL)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`INSERT INTO events(name) VALUES($1)`, urlarg)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
