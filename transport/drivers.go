package transport

import (
	"encoding/json"
	"game/service"
	"net/http"

	"github.com/gorilla/mux"
)

// FinishOrder accepts requests from the driver to finish an order.
func FinishOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var order service.Orders
	order.DriversID = params["id"]
	answer, err := order.FinishOrder(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

// DeleteOrder accepts requests from the driver to delete an order.
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	var order service.Orders
	params := mux.Vars(r)
	order.DriversID = params["id"]
	answer, err := order.DeleteOrder(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
