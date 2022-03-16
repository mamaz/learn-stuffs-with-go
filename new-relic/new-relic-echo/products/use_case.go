package products

import (
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
)

type ProductUC struct {
	repo *ProductRepo
}

func NewProductUC(repo *ProductRepo) *ProductUC {
	return &ProductUC{
		repo: repo,
	}
}

func (puc ProductUC) GetAllProducts(context echo.Context) []Product {
	txn := nrecho.FromContext(context)
	defer txn.StartSegment("GetAllProducts").End()

	return puc.repo.FindAll(context)
}

func (puc ProductUC) GetProductById(productID string, context echo.Context) (Product, bool) {
	txn := nrecho.FromContext(context)
	defer txn.StartSegment("GetProductById").End()
	return puc.repo.FindById(productID, context)
}

func (puc ProductUC) Create(newProduct CreateProductRequest, context echo.Context) Product {
	txn := nrecho.FromContext(context)
	defer txn.StartSegment("CreateNewProduct").End()

	return puc.repo.Create(newProduct, context)
}
