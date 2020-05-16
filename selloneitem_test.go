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
