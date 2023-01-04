package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetStatusEndPoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(map[string]string{"status": "OK"})
}


func main() {
	fmt.Println("Starting the application")
	
	router := mux.NewRouter()
	router.HandleFunc("/", GetStatusEndPoint).Methods("GET")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}
}

