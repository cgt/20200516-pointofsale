package pointofsale

import "fmt"

type InMemoryCatalog struct {
	pricesInCentsByBarcode map[string]int
}

func (s InMemoryCatalog) FormattedPrice(barcode string) (string, bool) {
	if priceInCents, ok := s.pricesInCentsByBarcode[barcode]; ok {
		return formatCentPrice(priceInCents), true
	}
	return "", false
}

func formatCentPrice(priceInCents int) string {
	return fmt.Sprintf("$%.2f", float64(priceInCents)/100.0)
}
