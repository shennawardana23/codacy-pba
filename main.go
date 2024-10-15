package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Initialize MySQL database connection
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL database")

	// Set up Gin router
	r := gin.Default()

	// Vulnerable endpoint
	r.GET("/user/vulnerable", getVulnerableUser)

	// Secure endpoint
	r.GET("/user/secure", getSecureUser)

	// Run the server
	r.Run(":8080")
}

// getVulnerableUser is vulnerable to SQL injection
func getVulnerableUser(c *gin.Context) {
	username := c.Query("username")

	// Vulnerable SQL query
	query := fmt.Sprintf("SELECT id, username, email FROM users WHERE username = '%s'", username)

	var id int
	var name, email string
	err := db.QueryRow(query).Scan(&id, &name, &email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "username": name, "email": email})
}

// getSecureUser uses parameterized queries to prevent SQL injection
func getSecureUser(c *gin.Context) {
	username := c.Query("username")

	// Secure SQL query using parameterized statement
	query := "SELECT id, username, email FROM users WHERE username = ?"

	var id int
	var name, email string
	err := db.QueryRow(query, username).Scan(&id, &name, &email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "username": name, "email": email})
}
