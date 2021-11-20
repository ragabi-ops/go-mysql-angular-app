package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetDBConnection() *sql.DB {
	con, err := sql.Open("mysql", "root:password@tcp(mysql-server:3306)/sakila")
	if err != nil {
		log.Fatal("Unable to open connection to db", err)
	}

	con.SetMaxOpenConns(20)
	con.SetMaxIdleConns(20)
	con.SetConnMaxLifetime(time.Minute * 5)

	return con
}
