package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Dbconnect() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	db_user := os.Getenv("DB_USER_NAME")
	db_pass := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_DATABASE_NAME")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", db_user, db_pass, db_host, db_port, db_name)

	db, err := sql.Open("mysql", conn)

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Attempt to ping the database to check the connection
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // Proper error handling instead of panic in your application
	}

	fmt.Println("Connected to the database!")

}
