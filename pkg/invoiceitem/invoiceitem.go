package invoiceitem

type Item struct {
	id      uint
	product string
	value   float64
}

// New retutns a new Item

func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}
