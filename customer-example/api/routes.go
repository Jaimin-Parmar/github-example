package api

import (
	customerApipk "postgres-crud/api/customerapi"
	"postgres-crud/app"

	"github.com/gorilla/mux"
)

type API struct {
	App *app.App
}

// New creates a new api
func New(a *app.App) (api *API, err error) {
	api = &API{App: a}
	return api, nil
}

func (a *API) InitializeRoutes(router *mux.Router) {
	customerAPI := customerApipk.New(a.App)
	router.HandleFunc("/customer", customerAPI.GetAllCustomer).Methods("GET")
	router.HandleFunc("/customer", customerAPI.CreateCustomer).Methods("POST")
	router.HandleFunc("/customer/{id}", customerAPI.GetOneCustomer).Methods("GET")
	router.HandleFunc("/customer/{id}", customerAPI.UpdateOneCustomer).Methods("PUT")
	router.HandleFunc("/customer/{id}", customerAPI.DeleteOneCustomer).Methods("DELETE")
}
