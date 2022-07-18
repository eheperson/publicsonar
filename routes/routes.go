package routes

import(
	"github.com/gorilla/mux"
	"ehe.com/publicsonar/controllers"
)

var PublicSonarRoutes = func(router *mux.Router){
	router.HandleFunc("/api/classifier", controllers.Classifier).Methods("GET")
}