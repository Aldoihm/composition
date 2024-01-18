package customer

type Customer struct {
	name    string
	address string
	phone   string
}

//Método que retorna a un nuevo cliente

func New(name, address, phone string) Customer {
	return Customer{name: name, address: address, phone: phone}
}
