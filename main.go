package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func main() {
	http.HandleFunc("/employee", get_employee_json)
	fmt.Printf("API endpoint -> http://localhost:8081/employee")
	http.ListenAndServe(":8081", nil)
}

type Employee struct {
	Id       string `json:"id"`
	First      string `json:"first"`
	Last        string `json:"last"`
	Department  string `json:"department ,omitempty"`
}

var employees = map[string]Employee{
	"0000000000": Employee{Id: "101", First: "Akshitha", Last: "Meka", Department: "Data Analytics"},
	"0000000001": Employee{Id: "102", First: "Bhargavi", Last: "Laveti", Department: "Data Analytics"},
	"0000000002": Employee{Id: "103", First: "Keerthi", Last: "Dwevedi", Department: "Data Analyics"},
	"0000000003": Employee{Id: "104", First: "Vamsi", Last: "Karri", Department: "Data Analytics"},
}

func get_employees() []Employee {
	values := make([]Employee, len(employees))
	i := 0
	for _, emp := range employees {
		values[i] = emp
		i++
	}
	return values
}

func get_employee_json(w http.ResponseWriter, r *http.Request) {
	emps := get_employees()
	data, err := json.Marshal(emps)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}
