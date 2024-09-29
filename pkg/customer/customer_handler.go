package customer

import (
	"encoding/json"
	"go-groc-store/pkg/database"
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	logger    *slog.Logger
	dbService database.DbService
}

func CustomerNewHandler(logger *slog.Logger, dbService database.DbService) *CustomerHandler {
	return &CustomerHandler{
		logger:    logger,
		dbService: dbService,
	}

}

func (handler *CustomerHandler) CreateCustomerTable() {
	handler.dbService.CreateCustomerTable()
}

func (handler *CustomerHandler) CreateCustomer(c *fiber.Ctx) error {
	newCustomer := new(Customer)

	if err := c.BodyParser(newCustomer); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "error": err})
	}

	jsonData, err := json.Marshal(newCustomer)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid customer object",
			"error":   err,
		})
	}

	customer, errr := handler.dbService.CreateCustomer(jsonData)

	if errr != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  400,
			"message": "Failed creating customer",
			"error":   errr,
		})
	}

	var id int64
	id, _ = customer.LastInsertId()

	return c.Status(200).JSON(fiber.Map{
		"status":  200,
		"message": "Create Successful",
		"id":      id,
	})

}

func (handler *CustomerHandler) GetAll(c *fiber.Ctx) error {
	customers := []Customer{}
	rows, err := handler.dbService.ListAllCustomers()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Customers not found", "data": nil})
	}

	defer rows.Close() // Ensure rows are closed after processing

	for rows.Next() {
		var customer Customer
		if err := rows.Scan(&customer.Id, &customer.CustomerId, &customer.FirstName,
			&customer.LastName, &customer.Email, &customer.Password, &customer.AddressOne, &customer.AddressTwo,
			&customer.City, &customer.State, &customer.Zip, &customer.PhoneNumber); err != nil {
			panic(err)
		}
		customers = append(customers, customer)
	}

	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "customers Found", "data": customers})

}

func (handler *CustomerHandler) GetCustomer(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	row := handler.dbService.GetCustomer(int64(id))
	var customer Customer
	if err := row.Scan(&customer.Id, &customer.CustomerId, &customer.FirstName,
		&customer.LastName, &customer.Email, &customer.Password, &customer.AddressOne, &customer.AddressTwo,
		&customer.City, &customer.State, &customer.Zip, &customer.PhoneNumber); err != nil {
		panic(err)
	}

	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "customer Found", "data": customer})

}
