package service

import (
	"context"
	"fmt"
	"game/repo"
)

// CreateCars makes a request to the database to create new car.
func (c Cars) CreateCars(ctx context.Context) (repo.Cars, error) {
	if c.ID == "" {
		return repo.Cars{}, fmt.Errorf("id should not be empty")
	}
	if c.DriversID != "" || c.DriversName != "" {
		if c.DriversID != "" {
			return repo.Cars{}, fmt.Errorf("cannot create the driver's name without the driver's ID")
		}
		if c.DriversName != "" {
			return repo.Cars{}, fmt.Errorf("cannot create the driver's ID  without the driver's name")
		}
		err := isDriver(ctx, c.DriversID)
		if err != nil {
			return repo.Cars{}, fmt.Errorf("Can't find driver. Error: %s", err.Error())
		}
	}
	res, err := repo.Cars(c).CreateCars(ctx)
	if err != nil {
		return repo.Cars{}, err
	}
	return res, nil
}

// GetCars makes a request to the database for a list of cars.
func (c Cars) GetCars(ctx context.Context) ([]repo.Cars, error) {
	res, err := repo.Cars(c).GetCars(ctx)
	if err != nil {
		return []repo.Cars{}, err
	}
	return res, nil
}

// DeleteCar makes a request to the database to delete car.
func (c Cars) DeleteCar(ctx context.Context) (string, error) {
	if c.ID == "" {
		return "", fmt.Errorf("id should not be empty")
	}
	err := repo.Cars(c).DeleteCar(ctx)
	if err != nil {
		return "", err
	}
	return "auto successfully deleted", nil
}

// UpdateCar makes a request to the database to update car.
func (c Cars) UpdateCar(ctx context.Context) (string, error) {
	if c.ID == "" {
		return "", fmt.Errorf("id should not be empty")
	}
	if c.DriversID != "" {
		err := isDriver(ctx, c.DriversID)
		if err != nil {
			return "", err
		}
	}
	err := repo.Cars(c).UpdateCar(ctx)
	if err != nil {
		return "", err
	}
	return "auto successfully updated", nil
}

// SetDriver makes a request to the database to set cars driver.
func (c Cars) SetDriver(ctx context.Context) (string, error) {
	if c.ID == "" {
		return "", fmt.Errorf("id should not be empty")
	}
	err := isDriver(ctx, c.DriversID)
	if err != nil {
		return "", fmt.Errorf("Can't find driver. Error: %s", err.Error())
	}
	err = repo.Cars(c).SetDriver(ctx)
	if err != nil {
		return "", err
	}
	return "the driver is successfully assigned to the car", nil
}

// UnsetDriver makes a request to the database to unset a cars driver.
func (c Cars) UnsetDriver(ctx context.Context) (string, error) {
	if c.ID == "" {
		return "", fmt.Errorf("id should not be empty")
	}
	var u Users
	u.ID, u.Name = "", ""
	err := repo.Cars(c).SetDriver(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Driver %s is successfully assigned to the car", c.DriversID), nil
}
