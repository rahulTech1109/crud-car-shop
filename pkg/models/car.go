package models

import (
	"muxcrud/pkg/config"

	"github.com/jinzhu/gorm"
)

var db * gorm.DB

type Car struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Company string `json:"company"`
	Variant string `json:"variant"`
}

func init(){
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Car{})
}

func (b *Car) CreateCar() *Car{
	db.NewRecord(b)
	db.Create(&b)
	return b 
}

func GetAllCar() []Car{
	var Cars []Car
	db.Find(&Cars)
	return Cars
}

func GetCarById(Id int64) (*Car , *gorm.DB){
	 var GetCar Car
	 db := db.Where("ID=?" , Id).Find(&GetCar)
	 return &GetCar, db
}

func CheckIfCarExists(name string)  (*Car, *gorm.DB){
      var car Car
	  db := db.Where("name=?", name).Find(&car)
	  return &car , db 
}

func DeleteCar(Id int64) bool {
	result := db.Delete(&Car{}, Id)
	return result.RowsAffected > 0
}