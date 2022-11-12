package main

import (
	"fmt"

	"github.com/Aldoihm/composition/pkg/customer"
	"github.com/Aldoihm/composition/pkg/invoice"
	"github.com/Aldoihm/composition/pkg/invoiceitem"
)

func main() {
	invoice := invoice.New(
		"México",
		"Teziutlán",
		customer.New("Aldo", "Av. Morelos", "2313228244"),
		invoiceitem.NewItems(
			invoiceitem.New(1, "papas sabritas adobadas", 12.34),
			invoiceitem.New(2, "chicle", 54.23),
			invoiceitem.New(3, "coca 1.750ml", 90.00),
		),
	)
	invoice.SetTotal()
	fmt.Printf("%+v", invoice)
}
