package pointofsale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type spyDisplay struct {
	currentText string
}

func (d *spyDisplay) Display(text string) {
	d.currentText = text
}

func TestSellOneItem(t *testing.T) {
	t.Run("product found", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display, nil}
		sale.OnBarcode("12345\n")

		assert.Equal(t, "$6.78", display.currentText)
	})

	t.Run("another product found", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display, nil}
		sale.OnBarcode("11223\n")

		assert.Equal(t, "$5.00", display.currentText)
	})

	t.Run("product found in catalog", func(t *testing.T) {
		display := &spyDisplay{}
		catalog := map[string]string{
			"55555": "$9.95",
		}

		sale := &Sale{display, catalog}
		sale.OnBarcode("55555\n")

		assert.Equal(t, "$9.95", display.currentText)
	})

	t.Run("product not found", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display, nil}
		sale.OnBarcode("::no such product::\n")

		assert.Equal(t, "product not found", display.currentText)
	})
}
