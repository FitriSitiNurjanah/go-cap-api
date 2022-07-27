package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct{
	//`` berfungsi untuk penamaan dalam json di postman
	ID int `json:"id xml:"id"`
	Name string `json:"name" xml:"name"` 
	City string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

var customers [] Customer = [] Customer{
	{1, "User1", "Jakarta", "12345"},
	{2, "User2", "Surabaya", "67890"},
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Celerates")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "get customer endpoint")

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
func getCustomer(w http.ResponseWriter, r *http.Request){

	// get route variable
	vars := mux.Vars(r)

	customerId := vars["customer_id"]

	//convert string to int
	id, err :=strconv.Atoi(customerId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid costumer id")
		return
	}

	//search customer data
	var cust Customer 

	for _, data := range customers{
		if data.ID == id {
			cust = data
		}
	}
	if cust.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "customer data not found")
		return
	}

	//return customer data
	w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(cust)
}

func addCustomer(w http.ResponseWriter, r * http.Request){
	//decode request body
	var cust Customer
	json.NewDecoder(r.Body).Decode(&cust)

	//generate new id
	nextID := getNextID()
	cust.ID = nextID

	// save data to array
	customers = append(customers, cust)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "customer successfully created")
}

func getNextID() int {
	cust := customers[len(customers)-1]

	return cust.ID + 1
}

// update customer
func updateCustomer(w http.ResponseWriter, r * http.Request){
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	id, err := strconv.Atoi(customerId)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid customer id")
		return
	}

	var cust Customer

	for customerIndex, data := range customers{
		if data.ID == id{
			cust = data

			var newCust Customer
			json.NewDecoder(r.Body).Decode(&newCust)

			customers[customerIndex].Name = newCust.Name
			customers[customerIndex].City = newCust.City
			customers[customerIndex].ZipCode = newCust.ZipCode

			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "customer data updated")
			return
		}
	}
	if cust.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "customer data not found")
		return
	}
	}


//delete customer
func deleteCustomer(w http.ResponseWriter, r * http.Request){
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	id, err := strconv.Atoi(customerId)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "invalid customer id")
		return
	}	

	for customerIndex, data := range customers{
		if data.ID == id {
			var newCust Customer
			json.NewDecoder(r.Body).Decode(&newCust)
			customers = append(customers[:customerIndex], customers[customerIndex+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "customer data deleted")
			return
		}
	}

}