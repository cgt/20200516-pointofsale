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
	if strings.TrimSpace(barcode) == "" {
		s.display.Display("error: invalid barcode")
		return
	}
	if price, ok := s.pricesByBarcode[strings.TrimSpace(barcode)]; ok {
		s.display.Display(price)
	} else {
		s.display.Display("product not found")
	}
}
