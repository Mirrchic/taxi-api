package transport

import (
	"encoding/json"
	"game/service"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAvailableCars gets a requests from a client to get a list of availeble cars.
func GetAvailableCars(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	AutoBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var Cars service.Cars
	err = json.Unmarshal(AutoBytes, &Cars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := Cars.GetAvailableCars(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
