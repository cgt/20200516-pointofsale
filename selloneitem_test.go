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

		sale := &Sale{display, map[string]string{
			"12345": "$6.78",
			"11223": "$5.00",
		}}
		sale.OnBarcode("12345\n")

		assert.Equal(t, "$6.78", display.currentText)
	})

	t.Run("another product found", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display, map[string]string{
			"12345": "$6.78",
			"11223": "$5.00",
		}}
		sale.OnBarcode("11223\n")

		assert.Equal(t, "$5.00", display.currentText)
	})

	t.Run("product not found", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display, map[string]string{
			"12345": "$6.78",
			"11223": "$5.00",
		}}
		sale.OnBarcode("::no such product::\n")

		assert.Equal(t, "product not found", display.currentText)
	})
}
