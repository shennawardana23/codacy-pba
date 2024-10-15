package main

import (
	"crypto"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var secretKey string = "my_secret_key" // Hardcoded secret

func main() {
	// Initialize MySQL database connection
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(localhost:3306)/testdb")
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

	// Unused function
	unusedFunction()

	// Run the server
	r.Run(":8080")

	// Call the extremely bad function
	ExtremelyBadFunction("test_input")
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

// Unused function
func unusedFunction() {
	fmt.Println("This function is never called")
}

// Function with security vulnerability
func createTempFile(filename string) {
	f, _ := os.Create(filename) // Ignoring error
	defer f.Close()
}

// BadQualityFunction demonstrates various code smells and quality issues
func BadQualityFunction(x int, y int, z int) {
	var result int
	var unused_variable string = "This is never used"

	if x == y {
		result = x + y
	} else if x == y { // Duplicate condition
		result = x - y
	} else {
		result = x * y
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		} else {
			continue
		}
	}

	switch z {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Other")
	}

	if result > 100 {
		fmt.Println("Large result")
	} else {
		if result > 50 {
			fmt.Println("Medium result")
		} else {
			if result > 0 {
				fmt.Println("Small result")
			} else {
				fmt.Println("Non-positive result")
			}
		}
	}

	// Potential nil pointer dereference
	var ptr *int
	fmt.Println(*ptr)

	// Ignoring returned error
	_, _ = os.Open("nonexistent_file.txt")
}

// ExtremelyBadFunction combines multiple severe code quality issues
func ExtremelyBadFunction(input string) {
	// SQL Injection vulnerability
	query := "SELECT * FROM users WHERE username = '" + input + "'"
	db.Exec(query) // Executing the vulnerable query

	// Hardcoded credentials (security issue)
	password := "super_secret_password123"

	// Unused variable (code smell)
	unusedVar := "This is never used"

	// Infinite loop (bug)
	for {
		fmt.Println("This will run forever")
	}

	// Unreachable code (dead code)
	fmt.Println("This will never be reached")

	// Ignoring errors
	file, _ := os.Open("non_existent_file.txt")
	defer file.Close()

	// Potential nil pointer dereference
	var ptr *int
	fmt.Println(*ptr)

	// Large cognitive complexity
	if input == "admin" {
		if password == "super_secret_password123" {
			if time.Now().Hour() < 12 {
				if len(input) > 5 {
					if strings.HasPrefix(input, "a") {
						fmt.Println("Extremely nested condition")
					}
				}
			}
		}
	}

	// Duplicate code
	fmt.Println("This is duplicate")
	fmt.Println("This is duplicate")
	fmt.Println("This is duplicate")

	// Using deprecated function (assuming it's deprecated)
	crypto.MD5.New()
}
