package tracing

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

func DBWithTransactionContext(gormdb *gorm.DB, echoContext echo.Context) *gorm.DB {
	transaction := nrecho.FromContext(echoContext)
	nrcontext := newrelic.NewContext(context.Background(), transaction)
	return gormdb.WithContext(nrcontext)
}
