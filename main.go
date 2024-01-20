package main

import (
	"fmt"

	"github.com/Aldoihm/composition/pkg/customer"
	"github.com/Aldoihm/composition/pkg/invoice"
	"github.com/Aldoihm/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Mexico",
		"Teziutlan",
		customer.New("Lupita", "Ixtlahuaca", "2311144378"),
		invoiceitem.NewItems(
			invoiceitem.New(1, "leche", 20.5),
			invoiceitem.New(2, "jamon", 25.8),
			invoiceitem.New(3, "queso", 15.2),
		),
	)
	i.SetTotal()
	fmt.Printf("%+v\n", i)
}
