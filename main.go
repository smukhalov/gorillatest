package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"gorillatest/routers"
)

func main() {
	router := routers.InitRoutes()

	router.Use(mux.CORSMethodMiddleware(router))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	handler := c.Handler(router)
	log.Println("Start server without CORS port 8000 ...")
	log.Fatal(http.ListenAndServe(":8000", handler))

	// srv := &http.Server{
	// 	Handler: router,
	// 	Addr:    "127.0.0.1:8000",
	// 	// Good practice: enforce timeouts for servers you create!
	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }

	// log.Fatal(srv.ListenAndServe())
}
