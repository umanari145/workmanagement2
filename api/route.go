package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//StartWebServer is entryPoint
func StartWebServer() error {
	fmt.Println("Rest API with Mux Routers")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/reserve", GetReserveItems).Methods("GET")
	//router.HandleFunc("/items", fetchAllItems).Methods("GET")
	//router.HandleFunc("/item/{id}", fetchSingleItem).Methods("GET")
	return http.ListenAndServe(fmt.Sprintf(":%d", 8080), router)
}
