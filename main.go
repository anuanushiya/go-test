// Application to test go is working are not
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func GetStatus(rw http.ResponseWriter, req *http.Request) {
	var result string
	result = "Running"
	rw.Write([]byte(result))
}

func GetVersion(rw http.ResponseWriter, req *http.Request) {
	version := "v2.1.1"
	rw.Write([]byte(version))
}

func main() {
	port := ":8700"

	route := mux.NewRouter()
	route.Headers("Content-Type", "application/json", "X-Requested-With", "XMLHttpRequest")

	/* Server Specific API */
	route.HandleFunc("/api/status", GetStatus).Methods("GET")
	route.HandleFunc("/api/version", GetVersion).Methods("GET")

	fmt.Println("Server listening at ", port)
	http.ListenAndServe(port, route)
}
