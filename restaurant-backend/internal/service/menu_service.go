package service

import (
	"context"
	"restaurant/internal/models"
	"restaurant/internal/repository"
)

type MenuApp struct {
	db repository.MenuDatabase
}

func NewMenuApp(db repository.MenuDatabase) *MenuApp {
	return &MenuApp{db: db}
}

func (app *MenuApp) GetAllMenu(ctx context.Context) ([]models.Menu, error) {
	return app.db.GetAllMenu(ctx)
}

func (app *MenuApp) GetMenuById(ctx context.Context, id int) (*models.Menu, error) {
	return app.db.GetMenuById(ctx, id)
}

func (app *MenuApp) AddMenu(ctx context.Context, menu *models.Menu) error {
	return app.db.AddMenu(ctx, menu)
}

func (app *MenuApp) UpdateMenu(ctx context.Context, id int, name *string, imageURL *string, desc *string, price *float64, category *string, available *bool) error {
	return app.db.UpdateMenu(ctx, id, name, imageURL, desc, price, category, available)
}

func (app *MenuApp) DeleteMenu(ctx context.Context, id int) error {
	return app.db.DeleteMenu(ctx, id)
}
