package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8084"

type Config struct {
	DB *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authenticatoin service")

	//connect to the database
conn := connectToDB()
if conn == nil {
	log.Panic("Can't connect to Postgres!")
}
	//setup config
	app := Config{
		DB: conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
} 