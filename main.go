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
		[]invoiceitem.Item{
			invoiceitem.New(1, "leche", 20),
			invoiceitem.New(2, "jamon", 25),
			invoiceitem.New(3, "queso", 15),
		},
	)
	i.SetTotal()
	fmt.Printf("%+v", i)
}
