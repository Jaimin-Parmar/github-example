package customer

import (
	"postgres-crud/app"
	"postgres-crud/model"

	"gorm.io/gorm"
)

type Service interface {
	CreateCustomer(customer model.Customer) error
	GetAllCustomer() ([]model.Customer, error)
	GetOneCustomer(id string) (model.Customer, error)
	UpdateOneCustomer(customer model.Customer, id string) error
	DeleteOneCustomer(id string) error
}

type service struct {
	DB *gorm.DB
}

func NewService(app *app.App) Service {
	svc := &service{
		DB: app.DB,
	}
	return svc
}

func (s *service) CreateCustomer(customer model.Customer) error {
	return createCustomer(s.DB, customer)
}

func (s *service) GetAllCustomer() ([]model.Customer, error) {
	return getAllCustomer(s.DB)
}

func (s *service) GetOneCustomer(id string) (model.Customer, error) {
	return getOneCustomer(s.DB, id)
}

func (s *service) UpdateOneCustomer(customer model.Customer, id string) error {
	return updateOneCustomer(s.DB, customer, id)
}

func (s *service) DeleteOneCustomer(id string) error {
	return deleteOneCustomer(s.DB, id)
}

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
