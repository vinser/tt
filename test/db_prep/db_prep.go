package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DSN string = "flibgo:flibgo@tcp(db:3306)/flibgo?charset=utf8"
)

type DB struct {
	*sql.DB
	Q map[string]*sql.Stmt
}

func main() {
	db := NewDB(DSN)
	defer db.Close()
	defer db.prepClose()

	start := time.Now()
	for i := 0; i < 10000; i++ {
		db.FindBook("прив")
	}
	finish := time.Now()
	elapsed := finish.Sub(start)
	fmt.Println("Time elapsed: ", elapsed)

	start = time.Now()
	db.addQuery("SELECT id FROM books WHERE sort LIKE ?")
	for i := 0; i < 10000; i++ {
		db.FindBookPrep("прив")
	}
	finish = time.Now()
	elapsed = finish.Sub(start)
	fmt.Println("Time elapsed: ", elapsed)

}

func NewDB(dsn string) *DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(10)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &DB{db, map[string]*sql.Stmt{}}
}

func (db *DB) FindBook(search string) int64 {
	var id int64 = 0
	q := "SELECT id FROM books WHERE sort LIKE ?"
	err := db.QueryRow(q, search).Scan(&id)
	if err == sql.ErrNoRows {
		return 0
	}
	return id
}

func (db *DB) FindBookPrep(search string) int64 {
	var id int64 = 0
	err := db.Q["SELECT id FROM books WHERE sort LIKE ?"].QueryRow(search).Scan(&id)
	if err == sql.ErrNoRows {
		return 0
	}
	return id
}

func (db *DB) addQuery(query string) {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	db.Q[query] = stmt
}

func (db *DB) prepClose() {
	for _, stmt := range db.Q {
		stmt.Close()
	}
}
