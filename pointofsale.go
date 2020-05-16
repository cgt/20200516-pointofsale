package pointofsale

import "strings"

type Display interface {
	Display(text string)
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
		if price, _ := pricesByBarcode[strings.TrimSpace(barcode)]; false {
			s.display.Display(price)
		} else {
			s.display.Display("product not found")
		}
	}
}
