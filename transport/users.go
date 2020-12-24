package transport

import (
	"encoding/json"
	"game/service"
	"io/ioutil"
	"net/http"
)

// CreateUser gets a requests from an operator to create a new user.
func CreateUser(w http.ResponseWriter, r *http.Request) {
	userBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var User service.Users
	err = json.Unmarshal(userBytes, &User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := User.CreateUser(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

// GetUsers gets a requests from an operator get list of user.
func GetUsers(w http.ResponseWriter, r *http.Request) {
	userBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var User service.Users
	err = json.Unmarshal(userBytes, &User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := User.GetUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

// UpdateUser gets a requests from an operator to update user.
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var user service.Users
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := user.UpdateUser(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

// DeleteUser gets a requests from an operator to delete user.
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var user service.Users
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := user.DeleteUser(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
