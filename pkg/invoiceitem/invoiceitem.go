package invoiceitem

type Item struct {
	id      uint
	product string
	value   float64
}

// New retutns a new Item

// Notar que en go no es necesario hacer esto {id: id,}
// solo con poner {id} go entiende que el valor que recibe se lo asigna a la estructura
func New(id uint, product string, value float64) Item {
	return Item{id, product, value}
}

// Getter of value
func (i Item) Value() float64 {
	return i.value
}
