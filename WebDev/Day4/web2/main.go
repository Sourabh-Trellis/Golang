package main

import (
	"fmt"
	"log"
	"net/http"
	"web2/middlewares"
	profilehandlers "web2/profileHandlers"
	userhandlers "web2/userHandlers"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title User Management API
// @version 1.0
// @description This is a sample server for managing users and their profiles.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()
	fmt.Println("http://localhost:8080")
	defer log.Fatal(http.ListenAndServe(":8080", r))

	//Middleware
	r.Use(middlewares.LoggingMiddleware)

	// Swagger route
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	//user routes
	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("", userhandlers.CreateUserHandler).Methods("POST")
	userRouter.HandleFunc("/{id}", userhandlers.GetUserHandler).Methods("GET")
	userRouter.HandleFunc("/search", userhandlers.QueryUserHandler).Methods("GET")

	//profile routes
	profileRouter := r.PathPrefix("/profile").Subrouter()
	profileRouter.HandleFunc("", profilehandlers.CreateProfileHandler).Methods("POST")
	profileRouter.HandleFunc("/{userID}", profilehandlers.GetProfileHandler).Methods("GET")
}
