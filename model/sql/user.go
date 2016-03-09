package model

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/slb350/axeco/shared/database"
)

// *****************************************************************************
// User
// *****************************************************************************

// User table contains the information for each user
type User struct {
	Id         uint32    `db:"id"`
	First_name string    `db:"first_name"`
	Last_name  string    `db:"last_name"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Status_id  uint8     `db:"status_id"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
	Deleted    uint8     `db:"deleted"`
}

// User_status table contains every possible user status (active/inactive)
type User_status struct {
	Id         uint8     `db:"id"`
	Status     string    `db:"status"`
	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
	Deleted    uint8     `db:"deleted"`
}

var (
	ErrCode        = errors.New("Case statement in code is not correct.")
	ErrNoResult    = errors.New("Result not found.")
	ErrUnavailable = errors.New("Database is unavailable.")
)

// Id returns the user id
func (u *User) ID() string {
	return fmt.Sprintf("%v", u.Id)
}

// standardizeErrors returns the same error regardless of the database used
func standardizeError(err error) error {
	if err == sql.ErrNoRows {
		return ErrNoResult
	}

	return err
}

// UserByEmail gets user information from email
func UserByEmail(email string) (User, error) {
	result := User{}
	err := database.Sql.Get(&result, "SELECT id, password, status_id, first_name FROM user WHERE email = ? LIMIT 1", email)
	return result, err
}

// UserIdByEmail gets user id from email
func UserIdByEmail(email string) (User, error) {
	result := User{}
	err := database.Sql.Get(&result, "SELECT id FROM user WHERE email = ? LIMIT 1", email)
	return result, err
}

// UserCreate creates user
func UserCreate(first_name, last_name, email, password string) error {
	_, err := database.Sql.Exec("INSERT INTO user (first_name, last_name, email, password) VALUES (?,?,?,?)", first_name, last_name, email, password)
	return err
}
