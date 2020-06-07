package routers

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = setUserRoutes(router)
	router = setSpaHandler(router)

	router.Use(loggingHandler)
	return router
}
