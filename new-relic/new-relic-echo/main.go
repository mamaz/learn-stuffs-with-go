package main

import (
	"log"
	"net/http"
	"new-relic-echo/products"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	StartServer()
}

func StartServer() {
	ec := echo.New()
	app := SetupNewRelic()

	ec.Use(nrecho.Middleware(app))

	// products
	repo := products.NewRepo()
	usecase := products.NewProductUC(repo)
	controller := products.NewController(usecase)

	ec.GET("/products", controller.GetAllProducts)
	ec.POST("/products", controller.CreateProduct)
	ec.GET("/products/:id", controller.GetProductById)
	ec.GET("/combined-products", controller.GetCombinedProducts)
	ec.POST("/error", controller.MakeError)
	ec.POST("/fatal", controller.MakeFatalError)
	ec.POST("/nullptr", controller.MakeNullPtr)

	if err := ec.Start(":9090"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func SetupNewRelic() *newrelic.Application {
	appName, IsAppNameSet := os.LookupEnv("APP_NAME")
	if !IsAppNameSet || appName == "" {
		log.Fatal("APP_NAME is not set")
	}

	licenseKey, IsLicenseSet := os.LookupEnv("LICENSE_KEY")
	if !IsLicenseSet || licenseKey == "" {
		log.Fatal("LICENSE_KEY is not set")
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(appName),
		newrelic.ConfigLicense(licenseKey),
		newrelic.ConfigDistributedTracerEnabled(true),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if err != nil {
		log.Fatalf("error establishing new relic %v", err.Error())
	}

	return app
}
