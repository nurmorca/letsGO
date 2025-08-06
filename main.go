package main

import (
	"context"
	"fmt"
	"letsgo/common/app"
	"letsgo/common/postgresql"
	"letsgo/controller"
	"letsgo/persistence"
	"letsgo/service"

	"github.com/labstack/echo/v4"
)

// web server is built with the help of echo library

func main() {
	ctx := context.Background()
	e := echo.New()
	configManager := app.NewConfigManager()

	dbPool := postgresql.GetConnectionPool(ctx, configManager.PostgreSqlConfig)
	productRepository := persistence.NewProductRepository(dbPool)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	productController.RegisterRoutes(e)
	for _, r := range e.Routes() {
		fmt.Printf("ROUTE: %s %s\n", r.Method, r.Path)
	}
	e.Start("localhost:8080")
}
