package invoiceitem

type Item struct {
	id      uint
	product string
	value   float64
}

func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

type Items []Item

// Método constructor para Items (Función variatica son los 3 puntitos)
func NewItems(items ...Item) Items {
	var is Items
	for _, item := range items {
		is = append(is, item)
	}
	return is
}

func (is Items) Total() float64 {
	var total float64 = 0
	for _, valor := range is {
		total += valor.value
	}
	return total
}
