package app

import (
	"fmt"
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

}
