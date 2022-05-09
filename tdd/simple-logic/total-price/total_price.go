package totalprice

import "github.com/shopspring/decimal"

type Items struct {
	ID       string
	Name     string
	Price    int
	Quantity int
}

const PRICE_FOR_DISCOUNT = 100000

func CountTotalPrice(items []Items, discount float64) decimal.Decimal {
	subtotal := calcSubtotal(items)

	if subtotal.GreaterThan(decimal.NewFromInt(PRICE_FOR_DISCOUNT)) {
		return discountedPrice(subtotal, discount)
	}

	return subtotal
}

func calcSubtotal(items []Items) decimal.Decimal {
	subtotal := decimal.NewFromInt(0)

	for _, item := range items {
		quantity := decimal.NewFromInt(int64(item.Quantity))
		price := decimal.NewFromInt(int64(item.Price))

		subtotal = subtotal.Add(quantity.Mul(price))
	}
	return subtotal
}

func discountedPrice(subtotal decimal.Decimal, discount float64) decimal.Decimal {
	percentageAfterDiscount := decimal.NewFromFloat(1.0).Sub(decimal.NewFromFloat(discount))
	discountedPrice := subtotal.Mul(percentageAfterDiscount)
	return discountedPrice
}
