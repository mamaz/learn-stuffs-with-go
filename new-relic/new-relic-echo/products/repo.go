package products

import (
	"errors"
	"math/rand"
	"time"

	"new-relic-echo/infrastructure/tracing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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

func (productRepo *ProductRepo) FindAll(echoContext echo.Context) ([]Product, error) {
	gormdb := tracing.DBWithTransactionContext(productRepo.gormdb, echoContext)

	var products []Product
	result := gormdb.Find(&products)
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Errorf("no data %v", result.Error)
		return []Product{}, nil
	}

	if result.Error != nil {
		log.Errorf("error on rerieving all data %v", result.Error)
		return nil, result.Error
	}

	return products, nil
}

func (productRepo *ProductRepo) FindById(productID string, echoContext echo.Context) (Product, bool) {
	gormdb := tracing.DBWithTransactionContext(productRepo.gormdb, echoContext)

	query := `
		select * from products where id = ?
	`
	var product Product
	time.Sleep(5 * time.Second) // simulate slowness so that it will be recorded on new relict tracing details
	db := gormdb.Raw(query, productID).Scan(&product)

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
	gormdb := tracing.DBWithTransactionContext(productRepo.gormdb, echoContext)

	newProduct := Product{
		ID:   uuid.NewString(),
		Name: product.Name,
		SKU:  product.SKU,
	}

	tx := gormdb.Create(&newProduct)

	if tx.Error != nil {
		return Product{}, tx.Error
	}

	return newProduct, nil
}
