package routes

import(
	"github.com/gorilla/mux"
	"ehe.com/publicsonar/controllers"
)

var PublicSonarRoutes = func(router *mux.Router){
	router.HandleFunc("/api/message-classifier", controllers.Classifier).Methods("GET")
	router.HandleFunc("/api/messages-json", controllers.MessagesJson).Methods("GET")
}