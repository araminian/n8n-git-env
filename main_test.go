package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestApp() *fiber.App {
	app := fiber.New()
	setupRoutes(app)
	return app
}

func TestGetProducts(t *testing.T) {
	app := setupTestApp()

	resp, err := app.Test(httptest.NewRequest("GET", "/products", nil))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var responseProducts []Product
	err = json.Unmarshal(body, &responseProducts)
	assert.NoError(t, err)
	assert.Len(t, responseProducts, 2)
}

func TestGetProduct(t *testing.T) {
	app := setupTestApp()

	// Test existing product
	resp, err := app.Test(httptest.NewRequest("GET", "/products/1", nil))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var product Product
	err = json.Unmarshal(body, &product)
	assert.NoError(t, err)
	assert.Equal(t, "1", product.ID)
	// Let's Break the test
	assert.Equal(t, "Laptop Broken", product.Name)
	//assert.Equal(t, "Laptop", product.Name)

	// Test non-existing product
	resp, err = app.Test(httptest.NewRequest("GET", "/products/999", nil))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestCreateProduct(t *testing.T) {
	app := setupTestApp()

	newProduct := Product{
		ID:    "3",
		Name:  "Tablet",
		Price: 299.99,
	}

	jsonBody, _ := json.Marshal(newProduct)
	req := httptest.NewRequest("POST", "/products", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	var createdProduct Product
	err = json.Unmarshal(body, &createdProduct)
	assert.NoError(t, err)
	assert.Equal(t, newProduct.ID, createdProduct.ID)
	assert.Equal(t, newProduct.Name, createdProduct.Name)
	//assert.Equal(t, newProduct.Price, createdProduct.Price)
	assert.Equal(t, newProduct.Price, createdProduct.Price+1.0)
}
