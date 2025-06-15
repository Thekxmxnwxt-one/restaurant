package repository

import (
	"context"
	"errors"
	"restaurant/internal/models"

	_ "github.com/lib/pq"
)

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	CreateEmployee(ctx context.Context, name, username, passwordHash string) (int, error)
	CreateCustomer(ctx context.Context, name, username, passwordHash string) (int, error)
}

func (pdb *PostgresDatabase) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User

	// ค้นหาจาก employee ก่อน
	row := pdb.db.QueryRowContext(ctx, `
		SELECT id, name, username, password_hash, 'staff' as role
		FROM employee WHERE username = $1 AND active = TRUE
	`, username)
	err := row.Scan(&user.ID, &user.Name, &user.Username, &user.PasswordHash, &user.Role)
	if err == nil {
		return &user, nil
	}

	// ถ้าไม่เจอใน employee ให้ค้นจาก customer
	row = pdb.db.QueryRowContext(ctx, `
		SELECT id, name, username, password_hash, 'customer' as role
		FROM customer WHERE username = $1
	`, username)
	err = row.Scan(&user.ID, &user.Name, &user.Username, &user.PasswordHash, &user.Role)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (pdb *PostgresDatabase) CreateEmployee(ctx context.Context, name, username, passwordHash string) (int, error) {
	var id int
	err := pdb.db.QueryRowContext(ctx, `
		INSERT INTO employee (name, username, password_hash)
		VALUES ($1, $2, $3) RETURNING id
	`, name, username, passwordHash).Scan(&id)
	return id, err
}

func (pdb *PostgresDatabase) CreateCustomer(ctx context.Context, name, username, passwordHash string) (int, error) {
	var id int
	err := pdb.db.QueryRowContext(ctx, `
		INSERT INTO customer (name, username, password_hash)
		VALUES ($1, $2, $3) RETURNING id
	`, name, username, passwordHash).Scan(&id)
	return id, err
}
