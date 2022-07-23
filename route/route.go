package route

import (
	"employee_project/api"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/employee", api.CreateEmployee).Methods("POST")
	r.HandleFunc("/employees", api.GetAllEmployees).Methods("GET")
	r.HandleFunc("/employees/{id}", api.GetEmployee).Methods("GET")
	r.HandleFunc("/employees/{id}", api.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employees/{id}", api.DeleteEmployee).Methods("DELETE")

	fmt.Println("Starting the server on localhost:8001")
	log.Fatal(http.ListenAndServe(":8001", r))
}
