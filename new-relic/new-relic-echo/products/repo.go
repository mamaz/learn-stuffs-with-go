package products

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

// Sleep within range [0,maxDuration), duration in milisceond
func sleepWithMax(maxDuration time.Duration) {
	randomMilisecond := time.Duration(rand.Intn(100))
	time.Sleep(randomMilisecond * time.Millisecond)
}

type ProductRepo struct {
	//id: Product
	db     map[string]Product
	gormdb *gorm.DB
}

func NewRepo(gormdb *gorm.DB) *ProductRepo {
	return &ProductRepo{
		gormdb: gormdb,
		db:     map[string]Product{},
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

func (productRepo *ProductRepo) FindById(productID string, echoContext echo.Context) (Product, bool) {
	transaction := nrecho.FromContext(echoContext)
	nrcontext := newrelic.NewContext(context.Background(), transaction)
	productRepo.gormdb = productRepo.gormdb.WithContext(nrcontext)

	query := `
		select * from products where id = ?
	`
	var product Product
	time.Sleep(5 * time.Second)
	db := productRepo.gormdb.Raw(query, productID).Scan(&product)

	if db.Error != nil && errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return Product{
			ID:   "",
			Name: "",
			SKU:  "",
		}, false
	}

	return product, true
}

func (productRepo *ProductRepo) Create(product CreateProductRequest, echoContext echo.Context) (Product, error) {
	transaction := nrecho.FromContext(echoContext)
	nrcontext := newrelic.NewContext(context.Background(), transaction)
	productRepo.gormdb = productRepo.gormdb.WithContext(nrcontext)

	newProduct := Product{
		ID:   uuid.NewString(),
		Name: product.Name,
		SKU:  product.SKU,
	}

	tx := productRepo.gormdb.Create(&newProduct)

	if tx.Error != nil {
		return Product{}, tx.Error
	}

	return newProduct, nil
}
