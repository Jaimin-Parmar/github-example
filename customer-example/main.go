package main

import (
	"fmt"
	"log"
	"net/http"
	"postgres-crud/api"
	"postgres-crud/app"

	"github.com/gorilla/mux"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(app)

	api, err := api.New(app)
	if err != nil {
		log.Fatalln(err)
	}

	serveAPI(api)
}

func serveAPI(api *api.API) {
	router := mux.NewRouter()
	api.InitializeRoutes(router)
	log.Println("server started at 8080")
	http.ListenAndServe(":8080", router)
}
