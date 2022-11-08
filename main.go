package main

import (
	"fmt"

	"github.com/Aldoihm/composition/pkg/customer"
	"github.com/Aldoihm/composition/pkg/invoice"
	"github.com/Aldoihm/composition/pkg/invoiceitem"
)

func main() {
	items := []invoiceitem.Item{
		invoiceitem.New(1, "papas sabritas adobadas", 10.5),
		invoiceitem.New(2, "chicle", 3.5),
		invoiceitem.New(3, "coca 1.750ml", 22),
	}
	invoice := invoice.New(
		"México",
		"Teziutlán",
		customer.New("Aldo", "Av. Morelos", "2313228244"),
		items,
	)
	invoice.SetTotal()
	fmt.Printf("%+v", invoice)
}
