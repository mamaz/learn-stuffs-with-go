package products

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func GetThirdParty(context echo.Context) ([]Product, error) {
	txn := nrecho.FromContext(context)

	client := &http.Client{}
	client.Transport = newrelic.NewRoundTripper(client.Transport)

	request, err := http.NewRequest("GET", "https://run.mocky.io/v3/75673fef-5f0f-4c6f-b223-7289f8fe18e3", nil)
	if err != nil {
		return nil, err
	}
	request = newrelic.RequestWithTransactionContext(request, txn)

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	products := []Product{}

	err = json.NewDecoder(response.Body).Decode(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
