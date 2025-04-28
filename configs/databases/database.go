package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
)

func ConnectionSqlServer() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseUser := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")

	log.Println("Database host: ", databaseHost)
	log.Println("Database port: ", databasePort)
	log.Println("Database user: ", databaseUser)
	log.Println("Database password: ", databasePassword)
	log.Println("Database name: ", databaseName)

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable;",
		databaseHost, databaseUser, databasePassword, databasePort, databaseName)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Ping failed:", err.Error())
		defer db.Close()
		return nil
	}

	fmt.Println("Connected!")

	return db
}
