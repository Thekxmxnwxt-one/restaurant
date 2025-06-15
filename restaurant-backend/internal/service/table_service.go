package service

import (
	"context"
	"restaurant/internal/models"
	"restaurant/internal/repository"
)

type TableApp struct {
	db repository.TableDatabase
}

func NewTableApp(db repository.TableDatabase) *TableApp {
	return &TableApp{db: db}
}

func (app *TableApp) GetAllTable(ctx context.Context) ([]models.Table, error) {
	return app.db.GetAllTable(ctx)
}

func (app *TableApp) GetTableById(ctx context.Context, id int) (*models.Table, error) {
	return app.db.GetTableById(ctx, id)
}

func (app *TableApp) AddTable(ctx context.Context, menu *models.Table) error {
	return app.db.AddTable(ctx, menu)
}

func (app *TableApp) UpdateTable(ctx context.Context, id int, name *string, capacity *int, status *string) error {
	return app.db.UpdateTable(ctx, id, name, capacity, status)
}

func (app *TableApp) DeleteTable(ctx context.Context, id int) error {
	return app.db.DeleteTable(ctx, id)
}
