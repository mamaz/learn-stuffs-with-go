package totalprice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalPrice(t *testing.T) {
	t.Run("it should have discount if total price is more than 100K IDR", func(t *testing.T) {
		// arrange
		items := []Items{
			{
				ID:       "123",
				Name:     "Bolu Kukus",
				Price:    15000,
				Quantity: 10,
			},
			{
				ID:       "123",
				Name:     "Brem Solo",
				Price:    10000,
				Quantity: 5,
			},
		}
		discounts := 0.1

		// act
		result := CountTotalPrice(items, discounts)

		// assert
		assert.Equal(t, result.StringFixedBank(2), "180000.00")
	})

	t.Run("it should not have discount if total price is less than 100K IDR", func(t *testing.T) {
		// arrange
		items := []Items{
			{
				ID:       "123",
				Name:     "Bala Bala",
				Price:    500,
				Quantity: 10,
			},
			{
				ID:       "12",
				Name:     "Cakwe",
				Price:    1000,
				Quantity: 5,
			},
		}
		discounts := 0.5

		// act
		result := CountTotalPrice(items, discounts)

		// assert
		assert.Equal(t, result.StringFixedBank(2), "10000.00")
	})

	t.Run("it should not have discount if total price is equal than 100K IDR", func(t *testing.T) {
		// arrange
		items := []Items{
			{
				ID:       "123",
				Name:     "Bala Bala",
				Price:    500,
				Quantity: 200,
			},
		}
		discounts := 0.5

		// act
		result := CountTotalPrice(items, discounts)

		// assert
		assert.Equal(t, result.StringFixedBank(2), "100000.00")
	})
}
