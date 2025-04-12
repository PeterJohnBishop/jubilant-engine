package main

import (
	"database/sql"
	"fmt"
	"jupilant-engine/main.go/pgdb"
	"jupilant-engine/main.go/server"
)

var db *sql.DB

func main() {
	fmt.Println("Jubilant-Engine, a Docker/Kubernetes playground.")

	db, err := pgdb.Connect(db)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	server.StartServer(db)
}
