package customer

// customer struct
type Customer struct {
	Id         int    `json:"id"`
	CustomerId string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`

	Email      string `json:"email"`
	Password   string `json:"password"`
	AddressOne string `json:"address_one"`

	AddressTwo string `json:"address_two"`
	City       string `json:"city"`
	State      string `json:"state"`

	Zip         int    `json:"zip"`
	PhoneNumber string `json:"phone_number"`
}

// Customers struct
type Customers struct {
	Customers []Customer `json:"customers"`
}
