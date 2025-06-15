package repository

import (
	"context"
	"database/sql"
	"fmt"
	"restaurant/internal/models"
	"strings"

	_ "github.com/lib/pq"
)

type TableDatabase interface {
	GetAllTable(ctx context.Context) ([]models.Table, error)
	GetTableById(ctx context.Context, id int) (*models.Table, error)
	AddTable(ctx context.Context, table *models.Table) error
	UpdateTable(ctx context.Context, id int, name *string, capacity *int, status *string) error
	DeleteTable(ctx context.Context, id int) error
}

func (pdb *PostgresDatabase) GetAllTable(ctx context.Context) ([]models.Table, error) {
	rows, err := pdb.db.QueryContext(ctx, "SELECT id, name, capacity, status FROM restaurant_table")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tables: %v", err)
	}
	defer rows.Close()

	var tables []models.Table
	for rows.Next() {
		var table models.Table

		if err := rows.Scan(&table.ID, &table.Name, &table.Capacity, &table.Status); err != nil {
			return nil, fmt.Errorf("failed to scan table: %v", err)
		}

		tables = append(tables, table)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed during row iteration: %v", err)
	}

	return tables, nil
}

func (pdb *PostgresDatabase) GetTableById(ctx context.Context, id int) (*models.Table, error) {
	query := "SELECT id, name, capacity, status FROM restaurant_table WHERE id = $1"
	row := pdb.db.QueryRowContext(ctx, query, id)

	var table models.Table
	err := row.Scan(&table.ID, &table.Name, &table.Capacity, &table.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no table found with id: %d", id)
		}
		return nil, fmt.Errorf("failed to get table by id: %v", err)
	}

	return &table, nil
}

func (pdb *PostgresDatabase) AddTable(ctx context.Context, table *models.Table) error {
	query := "INSERT INTO restaurant_table (name, capacity, status) VALUES ($1, $2, $3) RETURNING id"
	err := pdb.db.QueryRowContext(ctx, query,
		table.Name, table.Capacity, table.Status,
	).Scan(&table.ID)
	if err != nil {
		return fmt.Errorf("failed to add table: %v", err)
	}
	return nil
}

func (pdb *PostgresDatabase) UpdateTable(ctx context.Context, id int, name *string, capacity *int, status *string) error {
	setClauses := []string{}
	args := []interface{}{}
	argPos := 1

	if name != nil {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", argPos))
		args = append(args, *name)
		argPos++
	}
	if capacity != nil {
		setClauses = append(setClauses, fmt.Sprintf("capacity = $%d", argPos))
		args = append(args, *capacity)
		argPos++
	}
	if status != nil {
		setClauses = append(setClauses, fmt.Sprintf("status = $%d", argPos))
		args = append(args, *status)
		argPos++
	}

	if len(setClauses) == 0 {
		return fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf("UPDATE restaurant_table SET %s WHERE id = $%d",
		strings.Join(setClauses, ", "), argPos)
	args = append(args, id)

	_, err := pdb.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update table: %v", err)
	}
	return nil
}

func (pdb *PostgresDatabase) DeleteTable(ctx context.Context, id int) error {
	query := "DELETE FROM restaurant_table WHERE id = $1"
	result, err := pdb.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete table: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no table found with id: %d", id)
	}

	return nil
}
