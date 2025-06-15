package repository

import (
	"context"
	"database/sql"
	"fmt"
	"restaurant/internal/models"
	"strings"

	_ "github.com/lib/pq"
)

type MenuDatabase interface {
	GetAllMenu(ctx context.Context) ([]models.Menu, error)
	GetMenuById(ctx context.Context, id int) (*models.Menu, error)
	AddMenu(ctx context.Context, menu *models.Menu) error
	UpdateMenu(ctx context.Context, id int, name *string, imageURL *string, desc *string, price *float64, category *string, available *bool) error
	DeleteMenu(ctx context.Context, id int) error
}

func (pdb *PostgresDatabase) GetAllMenu(ctx context.Context) ([]models.Menu, error) {
	rows, err := pdb.db.QueryContext(ctx, "SELECT id, name, image_url, description, price, category, available FROM menu")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch menus: %v", err)
	}
	defer rows.Close()

	var menus []models.Menu
	for rows.Next() {
		var menu models.Menu

		if err := rows.Scan(&menu.ID, &menu.Name, &menu.ImageURL, &menu.Description, &menu.Price, &menu.Category, &menu.Available); err != nil {
			return nil, fmt.Errorf("failed to scan menu: %v", err)
		}

		menus = append(menus, menu)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during row iteration: %v", err)
	}

	return menus, nil
}

func (pdb *PostgresDatabase) GetMenuById(ctx context.Context, id int) (*models.Menu, error) {
	query := "SELECT id, name, image_url, description, price, category, available FROM menu WHERE id = $1"
	row := pdb.db.QueryRowContext(ctx, query, id)

	var menu models.Menu
	err := row.Scan(&menu.ID, &menu.Name, &menu.ImageURL, &menu.Description, &menu.Price, &menu.Category, &menu.Available)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no menu found with id: %d", id)
		}
		return nil, fmt.Errorf("failed to get menu by id: %v", err)
	}

	return &menu, nil
}

func (pdb *PostgresDatabase) AddMenu(ctx context.Context, menu *models.Menu) error {
	query := "INSERT INTO menu (name, image_url, description, price, category, available) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id"
	err := pdb.db.QueryRowContext(ctx, query,
		menu.Name, menu.ImageURL, menu.Description, menu.Price, menu.Category, menu.Available,
	).Scan(&menu.ID)
	if err != nil {
		return fmt.Errorf("failed to add menu: %v", err)
	}
	return nil
}

func (pdb *PostgresDatabase) UpdateMenu(ctx context.Context, id int, name *string, imageURL *string, desc *string, price *float64, category *string, available *bool) error {
	setClauses := []string{}
	args := []interface{}{}
	argPos := 1

	if name != nil {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", argPos))
		args = append(args, *name)
		argPos++
	}
	if imageURL != nil {
		setClauses = append(setClauses, fmt.Sprintf("image_url = $%d", argPos))
		args = append(args, *imageURL)
		argPos++
	}
	if desc != nil {
		setClauses = append(setClauses, fmt.Sprintf("description = $%d", argPos))
		args = append(args, *desc)
		argPos++
	}
	if price != nil {
		setClauses = append(setClauses, fmt.Sprintf("price = $%d", argPos))
		args = append(args, *price)
		argPos++
	}
	if category != nil {
		setClauses = append(setClauses, fmt.Sprintf("category = $%d", argPos))
		args = append(args, *category)
		argPos++
	}
	if available != nil {
		setClauses = append(setClauses, fmt.Sprintf("available = $%d", argPos))
		args = append(args, *available)
		argPos++
	}

	if len(setClauses) == 0 {
		return fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf("UPDATE menu SET %s WHERE id = $%d",
		strings.Join(setClauses, ", "), argPos)
	args = append(args, id)

	_, err := pdb.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update menu: %v", err)
	}
	return nil
}

func (pdb *PostgresDatabase) DeleteMenu(ctx context.Context, id int) error {
	query := "DELETE FROM menu WHERE id = $1"
	result, err := pdb.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete menu: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no menu found with id: %d", id)
	}

	return nil
}
