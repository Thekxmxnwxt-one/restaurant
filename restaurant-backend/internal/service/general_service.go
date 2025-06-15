package service

import (
	"restaurant/internal/repository"
)

type RestaurantApp struct {
	db repository.RestaurantDatabase
}

func NewRestaurantApp(db repository.RestaurantDatabase) *RestaurantApp {
	return &RestaurantApp{db: db}
}

func (app *RestaurantApp) Close() error {
	return app.db.Close()
}

func (app *RestaurantApp) Ping() error {
	return app.db.Ping()
}
