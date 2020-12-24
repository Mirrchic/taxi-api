package transport

import (
	"encoding/json"
	"game/service"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateOrder gets a requests from a clients to create a new order.
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	orderBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var order service.Orders
	err = json.Unmarshal(orderBytes, &order)
	params := mux.Vars(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := order.CreateOrder(r.Context(), params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

// GetAvailableOrders gets a requests from driver to get list of availeble orders .
func GetAvailableOrders(w http.ResponseWriter, r *http.Request) {
	var order service.Orders
	params := mux.Vars(r)
	answer, err := order.GetAvailebleOrders(r.Context(), params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

// UpdateOrder accepts requests from the driver to receive an order.
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	orderBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var order service.Orders
	err = json.Unmarshal(orderBytes, &order)
	params := mux.Vars(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := order.UpdateOrder(r.Context(), params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
