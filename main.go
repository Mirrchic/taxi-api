package main

import (
	"game/transport"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	//Handlers for operators cars control
	r.HandleFunc("/operator/cars/create", transport.CreateCars).Methods("POST")
	r.HandleFunc("/operator/cars/get", transport.GetCars).Methods("GET")
	r.HandleFunc("/operator/cars/delete", transport.DeleteCar).Methods("POST")
	r.HandleFunc("/operator/cars/update", transport.UpdateCar).Methods("POST")
	r.HandleFunc("/operator/cars/set", transport.SetDriver).Methods("POST")
	r.HandleFunc("/operator/cars/unset", transport.UnsetDriver).Methods("POST")

	//Handlers for operators user control
	r.HandleFunc("/operator/users/create", transport.CreateUser).Methods("POST")
	r.HandleFunc("/operator/users/get", transport.GetUsers).Methods("GET")
	r.HandleFunc("/operator/users/update", transport.UpdateUser).Methods("POST")
	r.HandleFunc("/operator/users/delete", transport.DeleteUser).Methods("POST")

	// Handlers for users requests.
	r.HandleFunc("/users/get_car/{id}", transport.GetAvailableCars).Methods("GET", "OPTIONS")
	r.HandleFunc("/users/orders/create/{id}", transport.CreateOrder).Methods("POST", "OPTIONS")

	// Handlers for drivers requests.
	r.HandleFunc("/drivers/orders/take/{id}", transport.UpdateOrder).Methods("POST", "OPTIONS")
	r.HandleFunc("/drivers/orders/availeble/{id}", transport.GetAvailableOrders).Methods("GET", "OPTIONS")
	r.HandleFunc("/drivers/orders/finish/{id}", transport.FinishOrder).Methods("GET", "OPTIONS")
	r.HandleFunc("/drivers/orders/delete/{id}", transport.DeleteOrder).Methods("GET", "OPTIONS")

	http.ListenAndServe(":", r)
}
