package repo

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1111"
	dbname   = "postgres"
)

// Orders is structure for working with orders data.
type Orders struct {
	ID          string `json:"id"`
	ClientsID   string `json:"clients_id"`
	ClientsName string `json:"clients_name"`
	Status      string `json:"status"`
	DriversID   string `json:"drivers_id"`
}

// Cars is structure for working with cars data.
type Cars struct {
	ID          string `json:"id"`
	Model       string `json:"model"`
	Coller      string `json:"coller"`
	DriversID   string `json:"drivers_id"`
	DriversName string `json:"drivers_name"`
}

// DBInit initiates access to the database
func DBInit() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
