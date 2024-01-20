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

// Declaración de tipo en base a un slice de estructura
type Items []Item

// Funcion constructora del nuevo tipo
func NewItems(items ...Item) Items {
	var is Items
	for _, item := range items {
		is = append(is, item)
	}
	return is
}

// Metodos del nuevo tipo
func (i Items) Total() float64 {
	var total float64
	for _, item := range i {
		total += item.value
	}
	return total
}
