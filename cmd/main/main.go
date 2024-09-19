package main

import (
	"log"
	"muxcrud/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() { 
      r := mux.NewRouter()
	  routes.RegisterCarRoutes(r)
	  http.Handle("/" , r )
	  log.Fatal(http.ListenAndServe("localhost:9010" , r))
}