package repo

import "context"

const createCar = `
INSERT INTO Cars (id, model, coller, drivers_id, drivers_name)
VALUES ($1, $2, $3, $4, $5) RETURNING id, model, coller, drivers_id, drivers_name
`

// CreateCars creates a new car in the database
func (arg Cars) CreateCars(ctx context.Context) (Cars, error) {
	db, err := DBInit()
	if err != nil {
		return Cars{}, err
	}

	row := db.QueryRowContext(ctx, createCar,
		arg.ID,
		arg.Model,
		arg.Coller,
		arg.DriversID,
		arg.DriversName,
	)
	var i Cars
	err = row.Scan(
		&i.ID,
		&i.Model,
		&i.Coller,
		&i.DriversID,
		&i.DriversName,
	)
	return i, err
}

const getCars = `
SELECT id, model, coller, drivers_id, drivers_name FROM Cars
`

// GetCars get a list of cars from the database.
func (arg Cars) GetCars(ctx context.Context) ([]Cars, error) {
	db, err := DBInit()
	if err != nil {
		return []Cars{}, err
	}
	rows, err := db.QueryContext(ctx, getCars)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Cars
	for rows.Next() {
		var i Cars
		if err := rows.Scan(
			&i.ID,
			&i.Model,
			&i.Coller,
			&i.DriversID,
			&i.DriversName,
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

const getAvailableCar = `
SELECT id, model, coller, drivers_id, drivers_name FROM Cars
WHERE drivers_id IS NOT null;
`

// GetAvailableCars get a cars thats available for take an order from the database.
func (arg Cars) GetAvailableCars(ctx context.Context) ([]Cars, error) {
	db, err := DBInit()
	if err != nil {
		return []Cars{}, err
	}
	rows, err := db.QueryContext(ctx, getAvailableCar)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Cars
	for rows.Next() {
		var i Cars
		if err := rows.Scan(
			&i.ID,
			&i.Model,
			&i.Coller,
			&i.DriversID,
			&i.DriversName,
		); err != nil {
			return nil, err
		}
		if i.DriversID != "" {
			items = append(items, i)
		}
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteCar = `
DELETE FROM Cars 
WHERE id = $1
`

// DeleteCar delets a car from the database.
func (arg Cars) DeleteCar(ctx context.Context) error {
	db, err := DBInit()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, deleteCar, arg.ID)
	return err
}

const updateCar = `
UPDATE  Cars SET model = $2, coller = $3, drivers_id = $4, drivers_name = $5
WHERE id = $1
`

// UpdateCar updates a cars data in the database.
func (arg Cars) UpdateCar(ctx context.Context) error {
	db, err := DBInit()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, updateCar,
		arg.ID,
		arg.Model,
		arg.Coller,
		arg.DriversID,
		arg.DriversName,
	)
	return err
}

const setDriver = `
UPDATE  Cars SET drivers_id = $2, drivers_name = $3
WHERE id = $1
`

// SetDriver set a cars driver in the database.
func (arg Cars) SetDriver(ctx context.Context) error {
	db, err := DBInit()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, setDriver,
		arg.ID,
		arg.DriversID,
		arg.DriversName,
	)
	return err
}
