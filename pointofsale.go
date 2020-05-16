package pointofsale

import (
	"fmt"
	"strings"
)

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

type InMemoryCatalog struct {
	formattedPricesByBarcode map[string]string
	pricesInCentsByBarcode   map[string]int
}

func (s InMemoryCatalog) FormattedPrice(barcode string) (string, bool) {
	var price string
	price, ok := s.formattedPricesByBarcode[barcode]
	if ok := ok; !ok {
		var priceInCents int
		priceInCents, ok = s.pricesInCentsByBarcode[barcode]
		price = fmt.Sprintf("$%v", float64(priceInCents)/100.0)
		return price, ok
	}
	return price, ok
}
