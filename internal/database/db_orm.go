package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	ID       = "ID"
	SHORTURL = "shorturl"
	REALURL  = "realurl"
)

var PG_USER string
var PG_PASSWORD string
var PG_DB_NAME string
var PG_DB_TABLE_NAME string

type DataRow struct {
	Id       int
	RealUrl  string
	ShortUrl string
}

func InsertRow(realUrl, shortUrl string, db *sql.DB) {
	q := fmt.Sprintf("insert into urls values(default, '%s', '%s')", realUrl, shortUrl)
	if _, err := db.Exec(q); err != nil {
		log.Println(err)
	}
}

func GetRow(column string, str string, db *sql.DB) *DataRow {
	q := fmt.Sprintf("select * from urls where %s = '%s'", column, str)
	d := new(DataRow)
	if err := db.QueryRow(q).Scan(&d.Id, &d.RealUrl, &d.ShortUrl); err != nil {
		return nil
	}
	return d
}

func createDBandTable() {
	connStr := fmt.Sprintf("user=%s password=%s sslmode=disable", PG_USER, PG_PASSWORD)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	command := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname = '%s'", PG_DB_NAME)
	response, err := db.Exec(command)
	if err != nil {
		log.Fatal(err)
	}
	if rows, _ := response.RowsAffected(); rows == 0 {
		command = fmt.Sprintf("CREATE DATABASE %s", PG_DB_NAME)
		_, err = db.Exec(command)
		if err != nil {
			log.Fatal(err)
		}
	}
	db.Close()
	command = fmt.Sprintf("%s dbname=%s", connStr, PG_DB_NAME)
	db, err = sql.Open("postgres", command)
	if err != nil {
		panic(err)
	}
	createBbTable := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id SERIAL PRIMARY KEY, %s TEXT UNIQUE, %s CHAR(10) UNIQUE)", PG_DB_TABLE_NAME, REALURL, SHORTURL)
	if _, err = db.Exec(createBbTable); err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func OpenDB() *sql.DB {
	connStr := fmt.Sprintf("user=%s password=%s sslmode=disable dbname=%s", PG_USER, PG_PASSWORD, PG_DB_NAME)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func InitDB() {
	createDBandTable()
}

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	err = godotenv.Load(filepath.Join(wd, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PG_USER = os.Getenv("PG_USER")
	PG_PASSWORD = os.Getenv("PG_PASSWORD")
	PG_DB_NAME = os.Getenv("PG_DB_NAME")
	PG_DB_TABLE_NAME = os.Getenv("PG_DB_TABLE_NAME")
}
