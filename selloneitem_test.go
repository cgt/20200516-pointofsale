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
	s.display.Display("$6.78")
}

func TestSellOneItem(t *testing.T) {
	t.Run("product found", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display}
		sale.OnBarcode("12345\n")

		assert.Equal(t, "$6.78", display.currentText)
	})
}
