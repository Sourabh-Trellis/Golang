package userhandlers

import (
	"encoding/json"
	"net/http"
	"web2/models"

	"github.com/gorilla/mux"
	
)

var users = make(map[string]models.User)

// var profiles = make(map[string]models.Profile)

// CreateUserHandler handles creating new user
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body User true "User"
// @Success 201 {object} User
// @Failure 400 {string} string "Bad Request"
// @Router /users [post]
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users[user.ID] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetUserHandler handles retrieving a user by ID
// @Summary Get a user by ID
// @Description Get user details by user ID
// @Tags users
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} User
// @Failure 404 {string} string "Not Found"
// @Router /users/{id} [get]
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	user, exists := users[userID]
	if !exists {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// QueryUserHandler handles getting user by name
// @Summary Query users by name
// @Description Get a list of users by name
// @Tags users
// @Produce  json
// @Param name query string true "Name"
// @Success 200 {array} User
// @Router /users/search [get]
func QueryUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	var result []models.User

	for _, user := range users {
		if user.Name == name {
			result = append(result, user)
		}
	}
	json.NewEncoder(w).Encode(result)
}
