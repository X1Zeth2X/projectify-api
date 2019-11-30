package models

import (
	"log"

	"github.com/X1Zeth2X/projectify-api/db"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var getDB = db.GetDB()

// User data structure
type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`

	// Optional
	FullName string `json:"full_name,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
	Bio      string `json:"bio,omitempty"`

	// Relationships
	// Projects []Project
}

// Validate incoming data
func (user *User) Validate() error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Username,
			// Username is required.
			validation.Required,
			// Length of 4-15
			validation.Length(4, 15),
			// Alphanumeric
			is.Alphanumeric,
		),

		validation.Field(&user.Email,
			// Email is required.
			validation.Required,
			// Is an email
			is.Email,
		),

		validation.Field(&user.Password,
			// Password is required.
			validation.Required,
			// Max length
			validation.Length(8, 255),
		),

		validation.Field(&user.FullName,
			// Length of 2-50
			validation.Length(2, 50),
			// Is Alphabetical
			is.Alpha,
		),
	)
}

// Create a new user.
func (user *User) Create() {
	// Add returns.

	err := user.Validate()
	if err != nil {
		log.Printf("Error occured while validating: %v\n", err)
		return // Add returns
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	user.Password = string(hashedPassword)
	// Get the current DB session and create the user.
	getDB.Create(user)

	if user.ID <= 0 {
		return // u.Message(status, "message")
	}

	// Return new user
}

// Create update method
// Create delete method
