package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const dsn = "host=localhost user=sunnex password=sunnex dbname=employee_db port=5432 sslmode=disable TimeZone=Africa/Lagos"

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "sunnex"
	DB_PASSWORD = "sunnex"
	DB_NAME     = "employee_db"
)

type Employees struct {
	gorm.Model

	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Cannot establish connection")
	}
	var employees Employees
	json.NewDecoder(r.Body).Decode(&employees)
	DB.Create(&employees)
	json.NewEncoder(w).Encode(employees)
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Cannot establish connection")
	}

	w.Header().Set("Content-Type", "application/json")
	var employees []Employees
	DB.Find(&employees)
	json.NewEncoder(w).Encode(&employees)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Cannot establish connection")
	}
	params := mux.Vars(r)
	var employees Employees
	DB.First(&employees, params["id"])
	json.NewEncoder(w).Encode(&employees)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("Cannot connect to  DB")
	}
	params := mux.Vars(r)

	var employee Employees
	DB.First(&employee, params["id"])
	json.NewDecoder(r.Body).Decode(&employee)
	DB.Save(&employee)
	json.NewEncoder(w).Encode(employee)

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Cannot connect to DB")
	}

	params := mux.Vars(r)
	var employees Employees
	DB.Delete(&employees, params["id"])
	json.NewEncoder(w).Encode("The employee is Deleted Successfully")

}
