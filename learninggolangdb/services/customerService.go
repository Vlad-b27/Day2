package service

import (
	"fmt"
	"learninggolangdb/db"
	"learninggolangdb/entity"
)

func GetAllCustomers() {
	customers := []entity.Customer{}
	db.GetDB().Where("name=?", "Profi").Find(&customers)
	fmt.Println(customers)

}

func DeleteCustomers(name string) {
	customer := entity.Customer{}
	db.GetDB().Where("name=?", name).Delete(&customer)
	fmt.Println(customer)

}

func CreateCustomer(customer entity.Customer) {
	db.GetDB().Create(&customer)
	fmt.Println(customer)
}
