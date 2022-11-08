package invoice

import (
	"github.com/Aldoihm/composition/pkg/customer"
	"github.com/Aldoihm/composition/pkg/invoiceitem"
)

// Invoice is the struct of a invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	for _, valor := range i.items {
		i.total += valor.Value()
	}
}

// constructor, returns a new Invoice
func New(country, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{country: country, city: city, client: client, items: items}
}
