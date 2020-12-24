package service

import (
	"context"
	"fmt"
	"game/repo"
)

// CreateOrder makes a request to the database to create a new order.
func (o Orders) CreateOrder(ctx context.Context, id string) (repo.Orders, error) {
	if o.DriversID != "" {
		return repo.Orders{}, fmt.Errorf("Error: Only the driver can pick the order")
	}
	err := isClient(ctx, id)
	if err != nil {
		return repo.Orders{}, fmt.Errorf("Error: cannot find client. Error: %s", err.Error())
	}
	o.Status = "waiting"
	res, err := repo.Orders(o).CreateOrder(ctx)
	if err != nil {
		return repo.Orders{}, fmt.Errorf("Error: cannot create order. Error: %s", err.Error())
	}
	return res, nil
}

// GetAvailebleOrders makes a request to the database for a list of available orders.
func (o Orders) GetAvailebleOrders(ctx context.Context, id string) ([]repo.Orders, error) {
	if id == "" {
		return []repo.Orders{}, fmt.Errorf("ID should not be empty")
	}
	err := isDriver(ctx, id)
	if err != nil {
		return []repo.Orders{}, fmt.Errorf("Can't find drivers ID. Error: %s", err.Error())
	}
	res, err := repo.Orders(o).GetAvailebleOrders(ctx, "whaiting")
	if err != nil {
		return []repo.Orders{}, err
	}
	return res, nil
}

func isStatusCorrect(s string) bool {
	status := map[string]bool{
		"waiting":     true,
		"in progress": true,
		"finished":    true,
	}
	return status[s]
}
