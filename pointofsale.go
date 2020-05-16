package pointofsale

import "strings"

type Catalog interface {
	FormattedPrice(barcode string) (price string, ok bool)
}

type Display interface {
	Display(text string)
}

type Sale struct {
	display Display
	catalog Catalog
}

func (s *Sale) OnBarcode(barcode string) {
	barcode = strings.TrimSpace(barcode)
	if barcode == "" {
		s.display.Display("error: invalid barcode")
		return
	}
	if price, ok := s.catalog.FormattedPrice(barcode); ok {
		s.display.Display(price)
	} else {
		s.display.Display("product not found")
	}
}
