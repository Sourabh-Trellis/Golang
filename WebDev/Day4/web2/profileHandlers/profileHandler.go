package profilehandlers

import (
	"encoding/json"
	"net/http"
	"web2/models"

	"github.com/gorilla/mux"
)

var profiles = make(map[string]models.Profile)

// CreateProfileHandler handles creating new profile
// @Summary Create a new profile
// @Description Create a new profile with the input payload
// @Tags profiles
// @Accept  json
// @Produce  json
// @Param profile body Profile true "Profile"
// @Success 201 {object} Profile
// @Failure 400 {string} string "Bad Request"
// @Router /profiles [post]
func CreateProfileHandler(w http.ResponseWriter, r *http.Request) {
	var profile models.Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	profiles[profile.UserID] = profile
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}

// GetProfileHandler handles retrieving a profile by user ID
// @Summary Get a profile by user ID
// @Description Get profile details by user ID
// @Tags profiles
// @Produce  json
// @Param userID path string true "User ID"
// @Success 200 {object} Profile
// @Failure 404 {string} string "Not Found"
// @Router /profiles/{userID} [get]
func GetProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]
	profile, exists := profiles[userID]
	if !exists {
		http.Error(w, "Profile not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(profile)
}
