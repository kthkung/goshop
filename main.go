package main

import (
	"fmt"

	"gorm.io/gorm"

	"go_shop/config"
	"go_shop/controller"

	"github.com/gofiber/fiber/v2"
)

func seedDatabase(db *gorm.DB) {
	db.AutoMigrate(
		&controller.Product{},
		&controller.Order{},
	)
	fmt.Println("Database created")
}

func setupRoutes(app *fiber.App) {

	product := app.Group("/api/products")
	product.Get("/", controller.GetProductList)
	product.Get("/:id", controller.GetProductById)
	product.Post("/", controller.AddProduct)

	order := app.Group("/api/admin/orders")
	order.Get("/", controller.GetOrderList)
	order.Post("/", controller.AddOrder)

}

func main() {
	app := fiber.New()

	config.DBConn = config.DatabaseConnection()
	seedDatabase(config.DBConn)
	setupRoutes(app)
	app.Listen(":8080")

	defer config.Close()
}
