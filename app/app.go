package app

import (
	"github.com/gorilla/mux"
	"github.com/ricardomaricato/banking/domain"
	"github.com/ricardomaricato/banking/service"
	"log"
	"net/http"
)

func Start() {

	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.CustomerRepositoryStub{})}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
