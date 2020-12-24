package service

// Cars is structure for working with cars data.
type Cars struct {
	ID          string `json:"id"`
	Model       string `json:"model"`
	Coller      string `json:"coller"`
	DriversID   string `json:"drivers_id"`
	DriversName string `json:"drivers_name"`
}

// Users is structure for working with users data.
type Users struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

// Orders is structure for working with  Orders data.
type Orders struct {
	ID          string `json:"id"`
	ClientsID   string `json:"clients_id"`
	ClientsName string `json:"clients_name"`
	Status      string `json:"status"`
	DriversID   string `json:"drivers_id"`
}
