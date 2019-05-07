package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/joho/godotenv"
	"database/sql"
)
import _ "github.com/go-sql-driver/mysql"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	isDebug := os.Getenv("IS_DEBUG")
	fmt.Println(isDebug)
	fmt.Println("START CONNECT TO MYSQL...")
	var builder strings.Builder
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	builder.WriteString(databaseUser)
	builder.WriteString(":")
	builder.WriteString(databasePassword)
	builder.WriteString("@tcp(")
	databaseIP := os.Getenv("DATABASE_IP")
	builder.WriteString(databaseIP)
	builder.WriteString(":3306)/")
	databaseName := os.Getenv("DATABASE_NAME")
	builder.WriteString(databaseName)
	var url = builder.String()
	fmt.Println(url)
	db, err := sql.Open("mysql", url)
	if err != nil {
		log.Fatal(err.Error());
		log.Fatal("Error connect to mysql")
	}
	defer db.Close()
	stmtOut, err := db.Prepare("SELECT CmpnyEnglishName FROM BS_Company WHERE CmpnyID = ?")
    if err != nil {
		log.Fatal(err.Error());
	}
	defer stmtOut.Close()
	var companyName string
	err = stmtOut.QueryRow("1").Scan(&companyName)
	if err != nil {
		log.Fatal(err.Error());
	}
	fmt.Println(companyName)
	rows, err := db.Query("SELECT LoginID FROM BS_CompanyUser WHERE CmpnyID = ?", 1)
	if err != nil {
		log.Fatal(err.Error());
	}
	defer rows.Close()
	var loginID string
	for rows.Next() {
		err = rows.Scan(&loginID)
		if err != nil {
			log.Fatal(err.Error());
		}
		fmt.Println(loginID)
	}
	// TODO
}
