# New Relic POC

Basically this repo is used for testing and learning capabilities of New Relic in these area:

* Basic Monitoring: Throughput (Request per minutes), how much time to complete one's request
* Application Performance Monitoring (APM): What causing an app to be slow, What kind of tracing it supports
* Error Tracking: what kind of error it can track, how can we install its tracking functionality
* Alerting

in a Go project.

## Running

create a New Relic project

create a .env file by copying and replacing values in .env.example

run `source .env`

run `go run main.go`

Try one of these endpoints

```go

// showing basic endpoint tracing, only on slow endpoint
ec.GET("/products", controller.GetAllProducts)
ec.POST("/products", controller.CreateProduct)
ec.GET("/products/:id", controller.GetProductById)

// APM demo, showing request span
ec.GET("/combined-products", controller.GetCombinedProducts)

// error demo
ec.POST("/error", controller.MakeError) 
ec.POST("/fatal", controller.MakeFatalError)
ec.POST("/nullptr", controller.MakeNullPtr)
```

License

MIT
