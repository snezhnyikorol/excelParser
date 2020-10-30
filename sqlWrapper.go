package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type sqlWrapper struct {
	db *sql.DB
}

type product struct {
	model string
	company string
	price int
}

func connectDb(path string) sqlWrapper {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	} else {
		return sqlWrapper{db}
	}
}

func (s sqlWrapper) insertProduct(pr product) {
	_, err := s.db.Exec("insert into products (model, company, price) values ($1, $2, $3)",
		pr.model, pr.company, pr.price)
	if err != nil{
		panic(err)
	}
}

func (s sqlWrapper) close() {
	err := s.db.Close()
	if err != nil {
		panic(err)
	}
}