package products

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
)

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
	txn := nrecho.FromContext(context)
	defer txn.StartSegment("FindAll").End()

	p := []Product{}

	for _, product := range productRepo.db {
		p = append(p, product)
	}
	return p
}

func (productRepo *ProductRepo) FindById(productID string, context echo.Context) (Product, bool) {
	txn := nrecho.FromContext(context)
	defer txn.StartSegment("FindById").End()

	if product, ok := productRepo.db[productID]; ok {
		return product, true
	}

	return Product{
		ID:   "",
		Name: "",
		SKU:  "",
	}, false
}

func (productRepo *ProductRepo) Create(product CreateProductRequest, context echo.Context) Product {
	txn := nrecho.FromContext(context)
	defer txn.StartSegment("CreateNewProductDB").End()

	newProduct := Product{
		ID:   uuid.NewString(),
		Name: product.Name,
		SKU:  product.SKU,
	}

	productRepo.db[newProduct.ID] = newProduct

	// simulate db operations
	time.Sleep(2 * time.Second)

	return newProduct
}
