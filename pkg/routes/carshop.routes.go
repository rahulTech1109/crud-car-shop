package routes

import (
	"muxcrud/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterCarRoutes = func (router *mux.Router)  {
	router.HandleFunc("/car/", controllers.CreateCar).Methods("POST")
	router.HandleFunc("/car/", controllers.GetAllCar).Methods("GET")
	router.HandleFunc("/car/{carId}", controllers.GetCarById).Methods("GET")
	router.HandleFunc("/car/exist/{name}", controllers.CheckIfCarExists).Methods("GET")
	router.HandleFunc("/car/{carId}", controllers.UpdateCar).Methods("PUT")
	router.HandleFunc("/car/{carId}", controllers.DeleteCar).Methods("DELETE")
}