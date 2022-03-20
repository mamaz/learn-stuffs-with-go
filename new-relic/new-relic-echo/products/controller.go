package products

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	usecase *ProductUC
}

func NewController(usecase *ProductUC) *Controller {
	return &Controller{
		usecase: usecase,
	}
}

func (c Controller) CreateProduct(context echo.Context) error {

	newproduct := new(CreateProductRequest)

	if err := context.Bind(newproduct); err != nil {
		return context.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   "Bad Request",
			"message": fmt.Sprintf("%v", err.Error()),
		})
	}

	created := c.usecase.Create(*newproduct, context)

	return context.JSON(http.StatusCreated, created)
}

func (c Controller) GetAllProducts(context echo.Context) error {
	products := c.usecase.GetAllProducts(context)

	return context.JSON(http.StatusOK, map[string]interface{}{
		"data": products,
	})
}

func (c Controller) GetProductById(context echo.Context) error {
	id := context.Param("id")

	product, found := c.usecase.GetProductById(id, context)

	if found {
		return context.JSON(http.StatusOK, map[string]interface{}{
			"data": product,
		})
	}

	return context.JSON(http.StatusOK, map[string]interface{}{
		"data": nil,
	})
}

func (c Controller) GetCombinedProducts(context echo.Context) error {
	products, err := c.usecase.GetCombinedProducts(context)
	if err != nil {
		return fmt.Errorf("error getting combined products %v", err)
	}
	return context.JSON(http.StatusOK, products)
}

func (c Controller) MakeError(context echo.Context) error {
	return context.JSON(http.StatusBadRequest, c.usecase.MakeError())
}

func (c Controller) MakeFatalError(context echo.Context) error {
	return context.JSON(http.StatusBadRequest, c.usecase.MakeFatalError())
}

func (c Controller) MakeNullPtr(context echo.Context) error {
	return context.JSON(http.StatusBadRequest, c.usecase.MakeNullPtr())
}
