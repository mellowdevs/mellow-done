package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	host, _ := os.LookupEnv("HOST")
	db_port, _ := os.LookupEnv("DB_PORT")
	user, _ := os.LookupEnv("DB_USER")
	password, _ := os.LookupEnv("DB_PASSWORD")
	dbname, _ := os.LookupEnv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, db_port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the db...")

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DB = db
	return DB
}

func GetDBInstance() *sql.DB {
	return DB
}
