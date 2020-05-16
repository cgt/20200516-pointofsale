package pointofsale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryCatalog(t *testing.T) {
	t.Run("lookup of barcode not in catalog fails", func(t *testing.T) {
		catalog := InMemoryCatalog{nil}
		_, ok := catalog.FormattedPrice("whatever")
		assert.False(t, ok)
	})

	t.Run("look up unformatted price by barcode then format it", func(t *testing.T) {
		catalog := InMemoryCatalog{
			map[string]int{"56789": 521},
		}
		price, ok := catalog.FormattedPrice("56789")
		assert.True(t, ok, "found price")
		assert.Equal(t, "$5.21", price)
	})

	t.Run("returns cent price if has both cent price and preformatted price", func(t *testing.T) {
		catalog := InMemoryCatalog{
			pricesInCentsByBarcode: map[string]int{"12345": 999},
		}
		price, ok := catalog.FormattedPrice("12345")
		assert.True(t, ok, "found price")
		assert.Equal(t, "$9.99", price)
	})

	t.Run("formats cent prices to two decimal places", func(t *testing.T) {
		catalog := InMemoryCatalog{
			pricesInCentsByBarcode: map[string]int{"12345": 100},
		}
		price, ok := catalog.FormattedPrice("12345")
		assert.True(t, ok, "found price")
		assert.Equal(t, "$1.00", price)
	})
}
