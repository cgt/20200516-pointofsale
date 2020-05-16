package pointofsale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryCatalog(t *testing.T) {
	t.Run("lookup of barcode not in catalog fails", func(t *testing.T) {
		catalog := InMemoryCatalog{nil, nil}
		_, ok := catalog.FormattedPrice("whatever")
		assert.False(t, ok)
	})

	t.Run("look up formatted price by barcode", func(t *testing.T) {
		catalog := InMemoryCatalog{map[string]string{"12345": "$6.78"}, nil}
		price, ok := catalog.FormattedPrice("12345")
		assert.True(t, ok, "found price")
		assert.Equal(t, "$6.78", price)
	})

	t.Run("look up unformatted price by barcode then format it", func(t *testing.T) {
		catalog := InMemoryCatalog{
			nil,
			map[string]int{"56789": 521},
		}
		price, ok := catalog.FormattedPrice("56789")
		assert.True(t, ok, "found price")
		assert.Equal(t, "$5.21", price)
	})
}
