package main

import (
	"fmt"

	"employee_project/api"
	"employee_project/route"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const dsn = "host=localhost user=sunnex password=sunnex dbname=employee_db port=5432 sslmode=disable TimeZone=Africa/Lagos"

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&api.Employees{})
}

func main() {
	InitialMigration()
	route.InitializeRouter()
}
