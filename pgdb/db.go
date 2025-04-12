package pgdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func LoadEnv() {
	// Get the absolute path to the .env file, relative to the current file
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	envPath := filepath.Join(basePath, ".env")
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file from %s", envPath)
	}
}

func Connect(db *sql.DB) (*sql.DB, error) {

	// load .env file
	LoadEnv()

	// set database connection parameters
	host := "postgres" // os.Getenv("PSQL_HOST")
	portStr := "5432"
	port, err := strconv.Atoi(portStr) // Convert to int
	if err != nil {
		return nil, err
	}
	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASSWORD")
	dbname := os.Getenv("PSQL_DBNAME")

	// connect to the database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println("Connecting to Postgres DB with connection string: ", psqlInfo)
	// create a new database instance to work with
	mydb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	// ping the server to test the connection
	err = mydb.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Successfully connected to Postgres DB: [%s] \n", dbname)
	return mydb, nil
}
