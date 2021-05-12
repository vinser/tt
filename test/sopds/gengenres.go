package main

import (
	"database/sql"
	"encoding/xml"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DSN string = "flibgo:flibgo@tcp(db:3306)/flibgo?charset=utf8"
)

type DB struct {
	*sql.DB
}

type RootDescr struct {
	Lang     string `xml:"lang,attr"`
	Title    string `xml:"genre-title,attr"`
	Detailed string `xml:"detailed,attr"`
}

type GenreAlt struct {
	Value  string `xml:"value,attr"`
	Format string `xml:"format,attr"`
}

type GenreDescr struct {
	Lang  string `xml:"lang,attr"`
	Title string `xml:"title,attr"`
}

type Subgenre struct {
	Value        string       `xml:"value,attr"`
	Descriptions []GenreDescr `xml:"genre-descr"`
	Alts         []GenreAlt   `xml:"genre-alt"`
}

type Genre struct {
	Value        string      `xml:"value,attr"`
	Descriptions []RootDescr `xml:"root-descr"`
	Subgenres    []Subgenre  `xml:"subgenres>subgenre"`
}

type GenresTree struct {
	XMLName xml.Name `xml:"fbgenrestransfer"`
	Genres  []Genre  `xml:"genre"`
}

type GenreRecord struct {
	ID    int
	Code  string
	Bunch string
	Name  string
}

func main() {
	db := NewDB(DSN)
	defer db.Close()
	makeGenresXML(db)
}

func makeGenresXML(db *DB) {

	q := `SELECT id, code, bunch, name FROM genres`
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	genres := []*GenreRecord{}

	for rows.Next() {
		g := &GenreRecord{}
		if err := rows.Scan(&g.ID, &g.Code, &g.Bunch, &g.Name); err != nil {
			log.Fatal(err)
		}
		genres = append(genres, g)
	}

}

func (db *DB) GenreBunches() ([]*GenreRecord, error) {
	q := `SELECT id, code, bunch, name FROM genres as g group by g.bunch`
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	genres := []*GenreRecord{}

	for rows.Next() {
		g := &GenreRecord{}
		if err := rows.Scan(&g.ID, &g.Code, &g.Bunch, &g.Name); err != nil {
			log.Fatal(err)
		}
		genres = append(genres, g)
	}
	return genres, nil
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
