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

func (puc ProductUC) GetAllProducts(context echo.Context) ([]Product, error) {
	txn := nrecho.FromContext(context)
	defer txn.StartSegment("GetAllProducts").End()

	return puc.repo.FindAll(context)
}

func (puc ProductUC) GetProductById(productID string, context echo.Context) (Product, bool) {
	return puc.repo.FindById(productID, context)
}

func (puc ProductUC) Create(newProduct CreateProductRequest, context echo.Context) (Product, error) {
	return puc.repo.Create(newProduct, context)
}

func (puc ProductUC) GetCombinedProducts(context echo.Context) ([]Product, error) {
	dbProducts, err := puc.GetAllProducts(context) // fetch from DB
	if err != nil {
		return nil, err
	}

	thirdPartyProducts, err := GetThirdParty(context) // fetch from 3rd party
	if err != nil {
		return nil, err
	}

	combined := append(dbProducts, thirdPartyProducts...)

	return combined, nil
}
