package handlers

import (
	"net/http"
)

// GetUser gets a specific user's info
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Get User ID
	// Do the usual db query stuff.
	// Return response with user info
}

// DeleteUser gets specific user and deletes it
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get User ID
	// Perform db stuff.
	// Return a response
}

// CreateUser gets incoming data and creates a user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Get incoming data, content n' stuff
	// Check if email/username is used.
	// Pass those data and create user.
	// Return new user and response
}

// UpdateUser gets incoming data and updates user data.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get incoming update data
	// Update user record
	// Return newly updated user and response
}
