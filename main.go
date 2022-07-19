package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"ehe.com/publicsonar/routes"
)

func main(){
	r := mux.NewRouter()
	routes.PublicSonarRoutes(r)
	http.Handle("/api/classifier", r)
	fmt.Print("Starting Server At Port : 8081\n")
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
