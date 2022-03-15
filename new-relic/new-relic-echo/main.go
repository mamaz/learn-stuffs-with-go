package main

import (
	"log"
	"net/http"
	"new-relic-echo/products"

	"github.com/labstack/echo/v4"
)

func main() {
	StartServer()
}

func StartServer() {
	ec := echo.New()

	repo := products.NewRepo()
	usecase := products.NewProductUC(repo)
	controller := products.NewController(usecase)

	ec.GET("/products", controller.GetAllProducts)
	ec.POST("/products", controller.CreateProduct)
	ec.GET("/products/:id", controller.GetProductById)

	if err := ec.Start(":9090"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
