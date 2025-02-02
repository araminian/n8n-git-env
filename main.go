package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var products = []Product{
	{ID: "1", Name: "Laptop", Price: 999.99},
	{ID: "2", Name: "Smartphone", Price: 499.99},
}

func setupRoutes(app *fiber.App) {
	app.Get("/products", getProducts)
	app.Get("/products/:id", getProduct)
	app.Post("/products", createProduct)
}

func getProducts(c *fiber.Ctx) error {
	return c.JSON(products)
}

func getProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	for _, product := range products {
		if product.ID == id {
			return c.JSON(product)
		}
	}
	return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
}

func createProduct(c *fiber.Ctx) error {
	product := new(Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid product data"})
	}

	products = append(products, *product)
	return c.Status(201).JSON(product)
}

func main() {
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
