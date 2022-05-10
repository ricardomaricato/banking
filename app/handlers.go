package app

import (
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipCode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ricardo", "Quatá", "19780-000"},
		{"Liliam", "Rancharia", "19760-000"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
	}
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Post request received")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	fmt.Println(id)
}
