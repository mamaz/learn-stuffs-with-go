package products

type Product struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

type CreateProductRequest struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
}
