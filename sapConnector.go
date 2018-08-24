package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

// SqlServer connection parameters
var server = "localhost"
var port = 1433
var user = "sa"
var password = "your_password"

// Global pointers for Gin router and Sql
var r *gin.Engine
var db *sql.DB

func main() {
	var err error

	// DataBase connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", server, user, password, port)

	// Create connection pool
	db, err = sql.Open("SqlServer", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Println("Connected to database")

	// Defer connection pool close after program exits
	defer db.Close()

	// TODO: Implement database pinging and version check

	// Router definition
	r = gin.Default()

	// Default status check route
	r.GET("/status", apiStatusCheck)

	// Initialize routes from routes.go
	initializeRoutes()

	// Run the router
	r.Run()

}
