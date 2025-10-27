package main

import (
	"net/http"

	"github.com/AxolotlJ-Dev/Api-Recipes/db"
	"github.com/AxolotlJ-Dev/Api-Recipes/models"
	"github.com/AxolotlJ-Dev/Api-Recipes/routes"
	"github.com/gorilla/mux"
)

func main() {

	db.DB.AutoMigrate(models.Recipes{}, models.User{})

	r := mux.NewRouter()

	r.Use(mux.CORSMethodMiddleware(r))

	r.HandleFunc("/", routes.HomeHandle)

	r.HandleFunc("/users", routes.GetUsersHandle).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandle).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandle).Methods("DELETE")

	// Recipes Routes

	r.HandleFunc("/recipes", routes.GetRecipesHandler).Methods("GET")
	r.HandleFunc("/recipes/{id}", routes.GetRecipeHandler).Methods("GET")
	r.HandleFunc("/recipes", routes.CreateRecipeHandler).Methods("POST", "OPTIONS")
	r.HandleFunc("/recipes/{id}", routes.DeleteRecipesHandler).Methods("Delete")

	// Login Routes

	r.HandleFunc("/login", routes.HomeLogin).Methods("GET")
	r.HandleFunc("/signIn", routes.GetLogin).Methods("POST", "OPTIONS")
	r.HandleFunc("/signUp", routes.PostUserHandle).Methods("POST", "OPTIONS")

	http.ListenAndServe(":8080", r)
}
