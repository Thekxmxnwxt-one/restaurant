package main

import (
	"context"
	"log"
	"time"

	"restaurant/internal/config"
	"restaurant/internal/handlers"
	"restaurant/internal/repository"
	"restaurant/internal/routes"
	"restaurant/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := repository.NewPostgresDatabase(cfg.GetConnectionString())
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}
	if db != nil {
		defer db.Close()
	}

	rs := service.NewRestaurantApp(db)
	h := handlers.NewRestaurantHandlers(rs)

	as := service.NewAuthApp(db)
	ah := handlers.NewAuthHandler(as)

	ms := service.NewMenuApp(db)
	mh := handlers.NewMenuHandlers(ms)

	ts := service.NewTableApp(db)
	th := handlers.NewTableHandlers(ts)

	go func() {
		for {
			time.Sleep(10 * time.Second)
			if err := db.Ping(); err != nil {
				log.Printf("Database connection lost: %v", err)

				if reconnErr := db.Reconnect(cfg.GetConnectionString()); reconnErr != nil {
					log.Printf("Failed to reconnect: %v", reconnErr)
				} else {
					log.Printf("Successfully reconnected to the database")
				}
			}
		}
	}()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:4200"},
		AllowMethods:  []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Authorization", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	r.Use(TimeoutMiddleware(5 * time.Second))

	r.GET("/health", h.HealthCheck)

	v1 := r.Group("/api/v1")
	routes.RegisterMenuRoutes(v1, mh)
	routes.RegisterTableRoutes(v1, th)
	routes.AuthRoutes(v1, ah)

	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Printf("Failed to run server: %v", err)
	}
}
