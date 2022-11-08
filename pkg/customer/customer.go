package customer

type Customer struct {
	name    string
	address string
	phone   string
}

// constructor de nuestro cliente
func New(name, address, phone string) Customer {
	return Customer{name, address, phone}
}
