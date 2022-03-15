package products

import "github.com/labstack/echo/v4"

type ProductUC struct {
	repo *ProductRepo
}

func NewProductUC(repo *ProductRepo) *ProductUC {
	return &ProductUC{
		repo: repo,
	}
}

func (puc ProductUC) GetAllProducts(context echo.Context) []Product {
	return puc.repo.FindAll()
}

func (puc ProductUC) GetProductById(context echo.Context, productID string) (Product, bool) {
	return puc.repo.FindById(productID)
}

func (puc ProductUC) Create(context echo.Context, newProduct CreateProductRequest) Product {
	return puc.repo.Create(newProduct)
}
