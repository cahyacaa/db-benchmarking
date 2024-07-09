package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // MariaDB driver
	_ "github.com/lib/pq"              // PostgreSQL driver
)

var (
	postgresDB *sql.DB
	mariaDB    *sql.DB
)

func init() {
	var err error
	postgresDB, err = sql.Open("postgres", "host=localhost user=postgres password=postgres dbname=benchmark_db sslmode=disable")
	if err != nil {
		panic(err)
	}

	dsn := "root:your_root_password@tcp(localhost:3306)/benchmark_db"
	mariaDB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

}
