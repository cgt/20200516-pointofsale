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
	catalog := map[string]string{
		"12345": "$6.78",
		"11223": "$5.00",
	}

	t.Run("product found", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display, catalog}
		sale.OnBarcode("12345\n")

		assert.Equal(t, "$6.78", display.currentText)
	})

	t.Run("another product found", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display, catalog}
		sale.OnBarcode("11223\n")

		assert.Equal(t, "$5.00", display.currentText)
	})

	t.Run("product not found", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display, catalog}
		sale.OnBarcode("::no such product::\n")

		assert.Equal(t, "product not found", display.currentText)
	})

	t.Run("empty barcode", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display, catalog}
		sale.OnBarcode("")

		assert.Equal(t, "error: invalid barcode", display.currentText)
	})

	t.Run("blank barcode", func(t *testing.T) {
		display := &spyDisplay{}

		sale := &Sale{display, catalog}
		sale.OnBarcode("\n")

		assert.Equal(t, "error: invalid barcode", display.currentText)
	})
}

func TestSellOneItemTable(t *testing.T) {
	testCases := []struct {
		description         string
		barcode             string
		expectedDisplayText string
	}{
		{
			description:         "product found",
			barcode:             "12345\n",
			expectedDisplayText: "$6.78",
		},
		{
			description:         "another product found",
			barcode:             "11223\n",
			expectedDisplayText: "$5.00",
		},
		{
			description:         "empty barcode",
			barcode:             "",
			expectedDisplayText: "error: invalid barcode",
		},
		{
			description:         "blank barcode",
			barcode:             "\n",
			expectedDisplayText: "error: invalid barcode",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			display := &spyDisplay{}
			catalog := map[string]string{
				"12345": "$6.78",
				"11223": "$5.00",
			}

			sale := &Sale{display, catalog}
			sale.OnBarcode(tc.barcode)

			assert.Equal(t, tc.expectedDisplayText, display.currentText)
		})
	}
}
