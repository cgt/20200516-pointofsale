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
}
