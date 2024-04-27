package main

import (
	"net/http"

	"github.com/AxolotlJ-Dev/Api-Recipes/db"
	"github.com/AxolotlJ-Dev/Api-Recipes/models"
	"github.com/AxolotlJ-Dev/Api-Recipes/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.Dbconnection()

	db.DB.AutoMigrate(models.Recipes{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandle)

	r.HandleFunc("/users", routes.GetUsersHandle).Methods("GET")
	r.HandleFunc("/users/{}", routes.GetUserHandle).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandle).Methods("POST")
	r.HandleFunc("/users", routes.DeleteUserHandle).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
