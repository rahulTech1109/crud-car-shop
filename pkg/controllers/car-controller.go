package controllers

import (
	"encoding/json"
	"fmt"
	"muxcrud/pkg/models"
	"muxcrud/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var NewCar models.Car

func GetAllCar(w http.ResponseWriter, r *http.Request){
	NewCar := models.GetAllCar()
	res , _ := json.Marshal(NewCar)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCarById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	carId := vars["carId"]
	ID, err := strconv.ParseInt(carId,0,0 )
	if err != nil {
		fmt.Println("get car by Id error while parsing")
	}
	carDetails, _:= models.GetCarById(ID)
	res ,_ :=json.Marshal(carDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCar(w http.ResponseWriter, r *http.Request){
	CreateCar := &models.Car{}
	utils.ParseBody(r , CreateCar)
	b := CreateCar.CreateCar()
	res , _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCar(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	carId := vars["carId"]
	ID , err := strconv.ParseInt(carId,0,0)
	if err != nil {
		fmt.Println("delete  error while parsing")
	} 
	car := models.DeleteCar(ID)
	res , _ := json.Marshal(car)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CheckIfCarExists(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "" {
		http.Error(w, "Car name is required", http.StatusBadRequest)
		return
	}

	car, result := models.CheckIfCarExists(name)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"message": "Car not found"})
		} else {
			http.Error(w, "Error checking for car: "+result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}

func UpdateCar(w http.ResponseWriter, r *http.Request){
	var UpdateCar = &models.Car{}
	utils.ParseBody(r, UpdateCar)
    vars := mux.Vars(r)
	carId := vars["carId"]
	ID , err := strconv.ParseInt(carId,0,0)
	if err != nil {
		fmt.Println("in update car error while parsing")
	}
	carDetails , db := models.GetCarById(ID)
	if UpdateCar.Name != " " {
	  carDetails.Name = UpdateCar.Name	
	}
	if UpdateCar.Variant != " " {
		carDetails.Variant = UpdateCar.Variant
	}
	if UpdateCar.Company != " "{
		carDetails.Company = UpdateCar.Company
	}
	db.Save(&carDetails)
	res, _ := json.Marshal(carDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

