package customerapi

import (
	"encoding/json"
	"log"
	"net/http"
	"postgres-crud/model"

	"github.com/gorilla/mux"
)

// Insert customer to database
func (a *api) CreateCustomer(w http.ResponseWriter, r *http.Request) {

	var customer model.Customer
	json.NewDecoder(r.Body).Decode(&customer)
	log.Println(customer)

	err := a.CustomerService.CreateCustomer(customer)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	log.Println("Customer data inserted successfully into database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// Get all customer from database
func (a *api) GetAllCustomer(w http.ResponseWriter, r *http.Request) {

	customers, err := a.CustomerService.GetAllCustomer()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (a *api) GetOneCustomer(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	customer, err := a.CustomerService.GetOneCustomer(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// Update one customer to database using id
func (a *api) UpdateOneCustomer(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	var customer model.Customer
	json.NewDecoder(r.Body).Decode(&customer)

	err := a.CustomerService.UpdateOneCustomer(customer, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	log.Println("Customer data updated successfully into database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}

// Delete one customer to database using id
func (a *api) DeleteOneCustomer(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	err := a.CustomerService.DeleteOneCustomer(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	log.Println("Customer data deleted successfully into database")

	w.Write([]byte("customer is deleted successfully with id " + id))
}
