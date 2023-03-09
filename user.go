package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var DB *gorm.DB
var err error

const DSN = "root:Ramayana@12345@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

func initialMigration() {
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to Database")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var users []User
	DB.Find(&users)
	json.NewEncoder(writer).Encode(users)

}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user User
	params := mux.Vars(request)
	DB.First(&user, params["id"])
	json.NewEncoder(writer).Encode(users)
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user User
	params := mux.Vars(request)
	DB.First(&user, params["id"])
	json.NewDecoder(request.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(writer).Encode(user)
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	DB.Create(&user)

	json.NewEncoder(writer).Encode(user)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var user User
	params := mux.Vars(request)
	DB.Delete(&user, params["id"])
	json.NewEncoder(writer).Encode("User is Deleted")
}
