package products

import "github.com/google/uuid"

type ProductRepo struct {
	//id: Product
	db map[string]Product
}

func NewRepo() *ProductRepo {
	return &ProductRepo{
		db: map[string]Product{},
	}
}

func (productRepo *ProductRepo) FindAll() []Product {
	p := []Product{}

	for _, product := range productRepo.db {
		p = append(p, product)
	}
	return p
}

func (productRepo *ProductRepo) FindById(productID string) (Product, bool) {
	if product, ok := productRepo.db[productID]; ok {
		return product, true
	}

	return Product{
		ID:   "",
		Name: "",
		SKU:  "",
	}, false
}

func (productRepo *ProductRepo) Create(product CreateProductRequest) Product {
	newProduct := Product{
		ID:   uuid.NewString(),
		Name: product.Name,
		SKU:  product.SKU,
	}

	productRepo.db[newProduct.ID] = newProduct

	return newProduct
}
