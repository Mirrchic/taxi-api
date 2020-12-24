package service

import (
	"context"
	"fmt"
	"game/repo"
)

// isDriver checks if role is driver
func isDriver(ctx context.Context, id string) error {
	var user repo.Users
	user.ID = id
	res, err := user.GetUser(ctx)
	if err != nil {
		return err
	}
	if res.Role != "driver" {
		return fmt.Errorf("Users with %s ID is not driver", id)
	}
	return nil
}

// UpdateOrder makes a request to the database to update an order.
func (o Orders) UpdateOrder(ctx context.Context, id string) (string, error) {
	if o.ID == "" {
		return "", fmt.Errorf("id should not be empty")
	}
	err := isDriver(ctx, id)
	if err != nil {
		return "", fmt.Errorf("Can't find drivers ID. Error: %s", err.Error())
	}
	err = checkDriversOrder(ctx, id, o.ID)
	if err != nil {
		return "", fmt.Errorf("Can't take order. Error: %s", err.Error())
	}
	o.Status = "in progress"

	err = repo.Orders(o).UpdateOrder(ctx, id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("order %s teken successfully", o.ID), nil
}

// checkDriversOrder checks if drivers have orrders.
func checkDriversOrder(ctx context.Context, id string, ridesID string) error {
	var order repo.Orders
	order.DriversID = id
	res, err := order.GetOrder(ctx)
	if err != nil {
		return err
	}
	if res.DriversID != "" {
		if res.ID == ridesID && res.Status == "in progress" {
			return nil
		}
		return fmt.Errorf("driver already have order")
	}
	return nil
}

// DeleteOrder makes a request to the database to delete an order.
func (o Orders) DeleteOrder(ctx context.Context) (string, error) {
	if o.DriversID == "" {
		return "", fmt.Errorf("id should not be empty")
	}
	order, err := repo.Orders(o).GetOrder(ctx)
	if err != nil {
		return "", err
	}
	if order.Status != "finished" {
		return "", fmt.Errorf("Error: order cannot be deleted until completed")
	}
	err = repo.Orders(o).DeleteOrder(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("order %s deleted successfully", o.ID), nil
}

// FinishOrder makes a request to the database to chage the order status to finished.
func (o Orders) FinishOrder(ctx context.Context) (string, error) {
	o.Status = "finished"
	err := isDriver(ctx, o.DriversID)
	if err != nil {
		return "", fmt.Errorf("Can't find drivers ID. Error: %s", err.Error())
	}
	err = repo.Orders(o).ChangeStatus(ctx)
	if err != nil {
		return "", fmt.Errorf("Error: Can't change status. Error: %s", err.Error())
	}
	return fmt.Sprintf("Order successfully finished"), nil
}
