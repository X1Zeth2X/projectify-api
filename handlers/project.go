package handlers

import (
	"net/http"
)

// GetProject gets a specific project.
func GetProject(w http.ResponseWriter, r *http.Request) {
	// Get item params
	// Perform get, db n' stuff.
	// render.JSON(w, r)
}

// CreateProject gets incoming data and creates a project.
func CreateProject(w http.ResponseWriter, r *http.Request) {
	// Get incoming data, content n' stuff
	// Pass those data and create em'
	// Return new project and response
}

// UpdateProject gets incoming data and updates project data.
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	// Get Project ID.
	// Perform db stuff.
	// Return a response.
}

// DeleteProject gets a specific project and deletes it.
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	// Get Project ID
	// Perform db delete.
	// Return a response
}

// GetProjects gets incoming id array of project IDs and return their info.
func GetProjects(w http.ResponseWriter, r *http.Request) {
	// Get IDs for projects
	// Grab those projects.
	// Return those cool projects and response
}
