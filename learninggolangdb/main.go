package main

import (
	"learninggolangdb/db"
	"learninggolangdb/entity"
	service "learninggolangdb/services"
)

func main() {
	db.InitDatabase()
	// service.GetAllCustomers()
	service.DeleteCustomers("Profi")
	customer := entity.Customer{
		Test:        "test",
		Address:     "test1",
		Description: "Test2",
	}
	service.CreateCustomer(customer)
}
