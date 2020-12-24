package repo

import (
	"context"
	"fmt"
)

const createOrder = `
INSERT INTO orders (id, clients_id, clients_name, status, drivers_id)
VALUES ($1, $2, $3, $4, $5) RETURNING id, clients_id, clients_name, status, drivers_id
`

// CreateOrder creates a new order in the database
func (arg Orders) CreateOrder(ctx context.Context) (Orders, error) {
	db, err := DBInit()
	if err != nil {
		return Orders{}, err
	}

	row := db.QueryRowContext(ctx, createOrder,
		arg.ID,
		arg.ClientsID,
		arg.ClientsName,
		arg.Status,
		arg.DriversID,
	)
	var i Orders
	err = row.Scan(
		&i.ID,
		&i.ClientsID,
		&i.ClientsName,
		&i.Status,
		&i.DriversID,
	)
	return i, err
}

const getAvailableOrders = `
SELECT id, clients_id, clients_name, status, drivers_id FROM orders
WHERE status = $1
`

// GetAvailebleOrders get an orders thats available for drivers from the database.
func (arg Orders) GetAvailebleOrders(ctx context.Context, s string) ([]Orders, error) {
	db, err := DBInit()
	if err != nil {
		return []Orders{}, err
	}
	rows, err := db.QueryContext(ctx, getAvailableOrders, s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Orders
	for rows.Next() {
		var i Orders
		if err := rows.Scan(
			&i.ID,
			&i.ClientsID,
			&i.ClientsName,
			&i.DriversID,
			&i.Status,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrder = `-- name: Order :exec
UPDATE orders SET Drivers_id = $2, Status = $3
WHERE id = $1
`

// UpdateOrder updates an order in the database.
func (arg Orders) UpdateOrder(ctx context.Context, id string) error {
	db, err := DBInit()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, updateOrder,
		arg.ID,
		arg.DriversID,
		arg.Status,
	)
	return err
}

const getOrder = `
SELECT * FROM orders
WHERE drivers_id  = $1
`

// GetOrder gets an order from the database.
func (arg Orders) GetOrder(ctx context.Context) (Orders, error) {
	db, err := DBInit()
	if err != nil {
		return Orders{}, err
	}
	rows := db.QueryRowContext(ctx, getOrder, arg.DriversID)

	var i Orders
	err = rows.Scan(
		&i.ID,
		&i.ClientsID,
		&i.ClientsName,
		&i.Status,
		&i.DriversID,
	)
	if err != nil {
		return Orders{}, fmt.Errorf("Can't get order. Error: %s", err.Error())
	}
	return i, nil
}

const deleteOrder = `
DELETE FROM orders
WHERE drivers_id  = $1
`

// DeleteOrder delets an order from the database.
func (arg Orders) DeleteOrder(ctx context.Context) error {
	db, err := DBInit()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, deleteOrder, arg.DriversID)
	return err
}

const finishOrder = `
UPDATE  orders SET status = $2
WHERE drivers_id = $1
`

// ChangeStatus changes the orderstatus in the database.
func (arg Orders) ChangeStatus(ctx context.Context) error {
	db, err := DBInit()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, finishOrder,
		arg.DriversID,
		arg.Status,
	)
	return err
}
