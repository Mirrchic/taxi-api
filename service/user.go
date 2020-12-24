package service

import (
	"context"
	"fmt"
	"game/repo"
)

// CreateUser makes a request to the database to create new user.
func (u Users) CreateUser(ctx context.Context) (repo.Users, error) {
	if !isRoleCorrect(u.Role) {
		return repo.Users{}, fmt.Errorf("Incorrect role: %s. The role must be either driver or user", u.Role)
	}
	res, err := repo.Users(u).CreateUser(ctx)
	if err != nil {
		return repo.Users{}, err
	}
	return res, nil
}

// GetUsers makes a request to the database to get a list of users.
func (u Users) GetUsers(ctx context.Context) ([]repo.Users, error) {
	res, err := repo.Users(u).GetUsers(ctx)
	if err != nil {
		return []repo.Users{}, err
	}
	return res, nil
}

// UpdateUser makes a request to the database to update a user.
func (u Users) UpdateUser(ctx context.Context) (string, error) {
	if u.ID == "" {
		return "", fmt.Errorf("id should not be empty")
	}
	if u.Role != "driver" && u.Role != "user" {
		return "", fmt.Errorf("Incorrect role: %s. The role must be either driver or user", u.Role)
	}
	err := repo.Users(u).UpdateUser(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("user %s updated successfully", u.ID), nil
}

// DeleteUser makes a request to the database to delete a user.
func (u Users) DeleteUser(ctx context.Context) (string, error) {
	if u.ID == "" {
		return "", fmt.Errorf("id should not be empty")
	}
	err := repo.Users(u).DeleteUser(ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("user %s deleted successfully", u.ID), nil
}

// isRoleCorrect checks if role is available.
func isRoleCorrect(r string) bool {
	roles := map[string]bool{
		"driver": true,
		"client": true,
	}
	return roles[r]
}
