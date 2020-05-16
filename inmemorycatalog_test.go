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
}
