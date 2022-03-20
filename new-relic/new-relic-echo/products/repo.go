package products

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Sleep within range [0,maxDuration), duration in milisceond
func sleepWithMax(maxDuration time.Duration) {
	randomMilisecond := time.Duration(rand.Intn(100))
	time.Sleep(randomMilisecond * time.Millisecond)
}

type ProductRepo struct {
	//id: Product
	db map[string]Product
}

func NewRepo() *ProductRepo {
	return &ProductRepo{
		db: map[string]Product{},
	}
}

func (productRepo *ProductRepo) FindAll(context echo.Context) []Product {
	p := []Product{}

	for _, product := range productRepo.db {
		p = append(p, product)
	}

	// simulate db ops
	sleepWithMax(200)

	return p
}

func (productRepo *ProductRepo) FindById(productID string, context echo.Context) (Product, bool) {
	if product, ok := productRepo.db[productID]; ok {
		return product, true
	}

	// simulate db ops
	sleepWithMax(100)

	return Product{
		ID:   "",
		Name: "",
		SKU:  "",
	}, false
}

func (productRepo *ProductRepo) Create(product CreateProductRequest, context echo.Context) Product {
	newProduct := Product{
		ID:   uuid.NewString(),
		Name: product.Name,
		SKU:  product.SKU,
	}

	productRepo.db[newProduct.ID] = newProduct

	// simulate db ops
	sleepWithMax(150)

	return newProduct
}
