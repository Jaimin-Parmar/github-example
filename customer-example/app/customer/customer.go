package customer

import (
	"postgres-crud/model"

	"gorm.io/gorm"
)

func createCustomer(db *gorm.DB, customer model.Customer) error {
	db.Create(&customer)
	return nil
}

func getAllCustomer(db *gorm.DB) ([]model.Customer, error) {
	var customers []model.Customer
	db.Find(&customers)
	return customers, nil
}

func getOneCustomer(db *gorm.DB, id string) (model.Customer, error) {
	var customer model.Customer
	db.First(&customer, id)
	return model.Customer{}, nil
}

func updateOneCustomer(db *gorm.DB, reqcustomer model.Customer, id string) error {
	dbcust, err := getOneCustomer(db, id)
	if err != nil {
		return err
	}

	dbcust.FirstName = reqcustomer.FirstName
	dbcust.LastName = reqcustomer.LastName
	dbcust.Email = reqcustomer.Email
	dbcust.Dateofbirth = reqcustomer.Dateofbirth
	dbcust.Mobilenumber = reqcustomer.Mobilenumber
	db.Save(&dbcust)
	return nil
}

func deleteOneCustomer(db *gorm.DB, id string) error {
	var customer model.Customer
	db.Delete(&customer, id)
	return nil
}
