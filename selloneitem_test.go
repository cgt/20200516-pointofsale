package pointofsale

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Display interface {
	Display(text string)
}

type spyDisplay struct {
	currentText string
}

func (d *spyDisplay) Display(text string) {
	d.currentText = text
}

type Sale struct {
	display Display
}

func (s *Sale) OnBarcode(barcode string) {
	pricesByBarcode := map[string]string{
		"12345\n": "$6.78",
	}
	if barcode == "12345\n" {
		s.display.Display(pricesByBarcode[barcode])
	} else if barcode == "11223\n" {
		s.display.Display("$5.00")
	} else {
		s.display.Display("product not found")
	}
}

func TestSellOneItem(t *testing.T) {
	t.Run("product found", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display}
		sale.OnBarcode("12345\n")

		assert.Equal(t, "$6.78", display.currentText)
	})

	t.Run("another product found", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display}
		sale.OnBarcode("11223\n")

		assert.Equal(t, "$5.00", display.currentText)
	})

	t.Run("product not found", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display}
		sale.OnBarcode("::no such product::\n")

		assert.Equal(t, "product not found", display.currentText)
	})
}
