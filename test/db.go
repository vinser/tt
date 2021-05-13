package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DSN         string = "flibgo:flibgo@tcp(db:3306)/flibgo?charset=utf8"
	INIT_SCRIPT string = "../database/mysql_db_init.sql"
	DROP_SCRIPT string = "../database/mysql_db_drop.sql"
)

type DB struct {
	*sql.DB
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("need action param: init or drop")
	} else {
		switch os.Args[1] {
		case "init":
			db := NewDB(DSN)
			defer db.Close()
			db.InitDB(INIT_SCRIPT)
		case "drop":
			db := NewDB(DSN)
			defer db.Close()
			db.DropDB(DROP_SCRIPT)
		default:
			fmt.Println("need action param: init or drop")
		}
	}

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
	return &DB{db}
}

func (db *DB) InitDB(initSQL string) {
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if !rows.Next() {
		db.execFile(initSQL)
	}
}

func (db *DB) DropDB(dropSQL string) {
	var err error
	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	if rows.Next() {
		db.execFile(dropSQL)
	}
}

func (db *DB) execFile(sqlFile string) {
	file, err := os.Open(sqlFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	q := ""

	for scanner.Scan() {
		q += scanner.Text()
		if strings.Contains(q, ";") {
			_, err := db.Exec(q)
			q = ""
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
