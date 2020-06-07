package routers

import (
	"gorillatest/controllers"

	"github.com/gorilla/mux"
)

func setUserRoutes(router *mux.Router) *mux.Router {
	//router.HandleFunc("/api/user", controllers.GetUser).Methods("GET")
	router.HandleFunc("/api/users/login", controllers.Login).Methods("POST")
	router.HandleFunc("/api/users", controllers.Register).Methods("POST")
	router.HandleFunc("/api/users/register", controllers.Register).Methods("POST")

	return router
}
