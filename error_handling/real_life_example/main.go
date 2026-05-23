package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// User model
type User struct {
	ID    int
	Email string
	Name  string
}

// User service handle database operations
type UserService struct {
	db *sql.DB
}

// Validation errors
type ValidationError struct {
	Field string
	Issue string
}

// Return validation error message
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on field '%s': %s", e.Field, e.Issue)
}

// User operation error
type UserError struct {
	Operation  string
	ErrorDetail error
}

// Return formatted user error
func (e *UserError) Error() string {
	return fmt.Sprintf("user operation '%s' failed: %v", e.Operation, e.ErrorDetail)
}

// Unwrap nested error
func (e *UserError) Unwrap() error {
	return e.ErrorDetail
}

// Setup database(SQLite) 
func setupDB() *sql.DB {

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	query := `
	CREATE TABLE users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE,
		name TEXT
	);
	`

	_, err = db.Exec(query)
	if err != nil {
		panic(err)
	}

	return db
}

// Check existing user(if email already exists)
func (s *UserService) checkUserExists(email string) (bool, error) {

	var count int

	err := s.db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Create user (validate and insert new user)
func (s *UserService) CreateUser(email, name string) error {

	// Validate email
	if !strings.Contains(email, "@") {
		return &UserError{
			Operation:  "create",
			ErrorDetail: &ValidationError{Field: "email", Issue: "invalid email format"},
		}
	}

	// Validate name
	if len(strings.TrimSpace(name)) < 2 {
		return &UserError{
			Operation:  "create",
			ErrorDetail: &ValidationError{Field: "name", Issue: "name too short"},
		}
	}

	// Check duplicate email
	exists, err := s.checkUserExists(email)
	if err != nil {
		return fmt.Errorf("failed checking user existence: %w", err)
	}

	if exists {
		return &UserError{
			Operation:  "create",
			ErrorDetail: fmt.Errorf("user with email '%s' already exists", email),
		}
	}

	// Insert user
	_, err = s.db.Exec("INSERT INTO users(email, name) VALUES(?, ?)", email, name)

	// error inserting user
	if err != nil {
		return fmt.Errorf("failed inserting user: %w", err)
	}

	return nil
}

// Get all users

// Retrieve all users
func (s *UserService) GetUsers() ([]User, error) {

	rows, err := s.db.Query(
		"SELECT id, email, name FROM users",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {

		var user User

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func main() {

	// Create service
	service := &UserService{
		db: setupDB(),
	}

	// Successful user creation
	err := service.CreateUser("john@example.com", "John")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User created successfully")
	}

	// Invalid email example
	err = service.CreateUser("yoo$example.yeah", "Mike")

	if err != nil {

		var userErr *UserError

		if errors.As(err, &userErr) {

			fmt.Println("\nUser operation error:")
			fmt.Println(userErr)

			var validationErr *ValidationError

			if errors.As(userErr.ErrorDetail, &validationErr) {

				fmt.Println("Validation details:")

				fmt.Printf("Field: %s\nIssue: %s\n", validationErr.Field, validationErr.Issue)
			}
		}
	}

	// Duplicate user example
	err = service.CreateUser("john@example.com", "Johnny")
	if err != nil {
		fmt.Println("\nDuplicate user error:")
		fmt.Println(err)
	}

	err = service.CreateUser("cornelius@example.com", "Cornelius")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User created successfully")
	}

	// Show all users
	users, err := service.GetUsers()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nUsers in database:")
	for _, user := range users {
		fmt.Printf("ID: %d | Email: %s | Name: %s\n", user.ID, user.Email, user.Name)
	}
}
