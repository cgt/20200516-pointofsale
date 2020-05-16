package pointofsale

import (
	"github.com/stretchr/testify/assert"
	"strings"
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
	display         Display
	pricesByBarcode map[string]string
}

func (s *Sale) OnBarcode(barcode string) {
	pricesByBarcode := map[string]string{
		"12345": "$6.78",
		"11223": "$5.00",
	}
	if price, ok := s.pricesByBarcode[strings.TrimSpace(barcode)]; ok {
		s.display.Display(price)
	} else {
		if price, ok := pricesByBarcode[strings.TrimSpace(barcode)]; ok {
			s.display.Display(price)
		} else {
			s.display.Display("product not found")
		}
	}
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
