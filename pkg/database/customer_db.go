package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/tidwall/gjson"
	"golang.org/x/crypto/bcrypt"
)

func (d *DbService) CreateCustomerTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS customers (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        customer_id TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		address_one TEXT NOT NULL,
		address_two TEXT NOT NULL,
		city TEXT NOT NULL,
		state TEXT NOT NULL,
		zip INTEGER NOT NULL,
		phone_number TEXT NOT NULL
    );`

	_, err := d.dbService.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
}

// CreateCustomer  (sql.Result, error)
func (d *DbService) CreateCustomer(customer []byte) (sql.Result, error) {

	fmt.Println(string(customer))
	firstName := gjson.Get(string(customer), "first_name")
	lastName := gjson.Get(string(customer), "last_name")
	customerId := gjson.Get(string(customer), "customer_id")
	email := gjson.Get(string(customer), "email")
	password := gjson.Get(string(customer), "password")
	addressOne := gjson.Get(string(customer), "address_one")
	addressTwo := gjson.Get(string(customer), "address_two")
	city := gjson.Get(string(customer), "city")
	state := gjson.Get(string(customer), "state")
	zip := gjson.Get(string(customer), "zip")
	phoneNumber := gjson.Get(string(customer), "phone_number")

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password.String()), 8)

	fmt.Println(string(hashedPassword))
	result, err := d.dbService.Exec("INSERT INTO customers (first_name, last_name, customer_id, email, password, address_one, address_two, city, state, zip, phone_number ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		firstName.String(), lastName.String(), customerId.String(), email.String(),
		hashedPassword, addressOne.String(), addressTwo.String(), city.String(),
		state.String(), zip.Int(), phoneNumber.String())

	return result, err
}

func (d *DbService) ListAllCustomers() (*sql.Rows, error) {
	return d.dbService.Query("SELECT * FROM  customers ORDER BY id LIMIT 300 OFFSET 0;")
}

func (d *DbService) GetCustomer(id int64) *sql.Row {
	return d.dbService.QueryRow("SELECT * FROM  customers where id =?", id)
}
