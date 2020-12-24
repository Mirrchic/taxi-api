package service

import (
	"context"
	"fmt"
	"game/repo"
)

// GetAvailableCars makes a request to the database to get a list of available cars.
func (c Cars) GetAvailableCars(ctx context.Context, clientID string) ([]repo.Cars, error) {
	err := isClient(ctx, clientID)
	if err != nil {
		return []repo.Cars{}, fmt.Errorf("Can't find clent. Error: %s", err.Error())
	}
	res, err := repo.Cars(c).GetAvailableCars(ctx)
	if err != nil {
		return []repo.Cars{}, err
	}
	return res, nil
}

func isClient(ctx context.Context, s string) error {
	var u repo.Users
	u.ID = s
	res, err := u.GetUser(ctx)
	if err != nil {
		return err
	}
	if res.Role != "client" {
		return fmt.Errorf("User with %s ID is not client", s)
	}
	return nil
}
