package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Pincode string `json:"pincode" xml:"pincode"`
}

type CurrentTime struct {
	Time string `json:"current_time"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Aditya", City: "Bangalore", Pincode: "560037"},
		{Name: "Sai", City: "Bangalore", Pincode: "560037"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "POST REQUESET RECIVED")
}

func getTime(w http.ResponseWriter, r *http.Request) {
	timezone := r.URL.Query().Get("tz")
	layout := "2006-01-02 15:04:05 -0700 MST"
	if timezone != "" {
		location, err := time.LoadLocation(timezone)
		if err != nil {
			fmt.Fprint(w, err)
		}
		currentTime := CurrentTime{time.Now().In(location).Format(layout)}
		if r.Header.Get("Content-Type") == "application/xml" {
			w.Header().Add("Content-Type", "application/xml")
			xml.NewEncoder(w).Encode(currentTime)
		} else {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(currentTime)
		}
	} else {
		currentTime := CurrentTime{time.Now().Format(layout)}
		if r.Header.Get("Content-Type") == "application/xml" {
			w.Header().Add("Content-Type", "application/xml")
			xml.NewEncoder(w).Encode(currentTime)
		} else {
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(currentTime)
		}
	}
}
