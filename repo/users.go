package repo

import (
	"context"
)

// Users is structure for working with users data.
type Users struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

const createUser = `
INSERT INTO users (id, name, role)
VALUES ($1, $2, $3) RETURNING id, name, role
`

// CreateUser creates a new user in the database
func (arg Users) CreateUser(ctx context.Context) (Users, error) {
	db, err := DBInit()
	if err != nil {
		return Users{}, err
	}
	row := db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Role,
	)
	var i Users
	err = row.Scan(
		&i.ID,
		&i.Name,
		&i.Role,
	)
	return i, err
}

const getUsers = `
SELECT * FROM users
`

// GetUsers get a list of users from the database.
func (arg Users) GetUsers(ctx context.Context) ([]Users, error) {
	db, err := DBInit()
	if err != nil {
		return []Users{}, err
	}
	rows, err := db.QueryContext(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Users
	for rows.Next() {
		var i Users
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Role,
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

const getUser = `
SELECT * FROM users
WHERE ID = $1
`

// GetUser get a user from the database.
func (arg Users) GetUser(ctx context.Context) (Users, error) {
	db, err := DBInit()
	if err != nil {
		return Users{}, err
	}
	rows := db.QueryRowContext(ctx, getUser, arg.ID)

	var i Users
	err = rows.Scan(
		&i.ID,
		&i.Name,
		&i.Role,
	)
	if err != nil {
		return Users{}, err
	}
	return i, nil
}

const deleteUser = `
DELETE FROM users 
WHERE id = $1
`

// DeleteUser delete a user from the database.
func (arg Users) DeleteUser(ctx context.Context) error {
	db, err := DBInit()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, deleteUser, arg.ID)
	return err
}

const updateUser = `
UPDATE  users SET Name = $2, Role = $3
WHERE id = $1
`

// UpdateUser updates a user in the database.
func (arg Users) UpdateUser(ctx context.Context) error {
	db, err := DBInit()
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Role,
	)
	return err
}
