package transport

import (
	"encoding/json"
	"game/service"
	"io/ioutil"
	"net/http"
)

// CreateCars gets a requests from an operator to create new car.
func CreateCars(w http.ResponseWriter, r *http.Request) {
	CarsBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var Cars service.Cars
	err = json.Unmarshal(CarsBytes, &Cars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := Cars.CreateCars(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

// GetCars gets a requests from an operator to get a list of cars.
func GetCars(w http.ResponseWriter, r *http.Request) {
	AutoBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var Cars service.Cars
	err = json.Unmarshal(AutoBytes, &Cars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := Cars.GetCars(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

// DeleteCar gets a requests from an operator to delete a car.
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	AutoBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var Car service.Cars
	err = json.Unmarshal(AutoBytes, &Car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := Car.DeleteCar(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

// UpdateCar gets a requests from an operator to update the cars information.
func UpdateCar(w http.ResponseWriter, r *http.Request) {
	autoBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var Car service.Cars
	err = json.Unmarshal(autoBytes, &Car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := Car.UpdateCar(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

// SetDriver gets a requests from an operator to  set a driver to a car.
func SetDriver(w http.ResponseWriter, r *http.Request) {
	autoBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var Car service.Cars
	err = json.Unmarshal(autoBytes, &Car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := Car.SetDriver(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}

// UnsetDriver gets a requests from an operator to  set a driver to a car.
func UnsetDriver(w http.ResponseWriter, r *http.Request) {
	autoBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var Car service.Cars
	err = json.Unmarshal(autoBytes, &Car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	answer, err := Car.UnsetDriver(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
