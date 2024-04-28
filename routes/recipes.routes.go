package routes

import (
	"encoding/json"
	"net/http"

	"github.com/AxolotlJ-Dev/Api-Recipes/db"
	"github.com/AxolotlJ-Dev/Api-Recipes/models"
	"github.com/gorilla/mux"
)

func GetRecipesHandler(w http.ResponseWriter, r *http.Request) {
	var recipes []models.Recipes
	db.DB.Find(&recipes)
	json.NewEncoder(w).Encode(recipes)
}

func CreateRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var recipes models.Recipes
	json.NewDecoder(r.Body).Decode(&recipes)

	createRecipe := db.DB.Create(&recipes)
	err := createRecipe.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&recipes)

}

func GetRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var recipes models.Recipes
	params := mux.Vars(r)

	db.DB.Find(&recipes, params["id"])

	if recipes.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("recipe not found"))
		return
	}

	json.NewEncoder(w).Encode(&recipes)

}

func DeleteRecipesHandler(w http.ResponseWriter, r *http.Request) {
	var recipes models.Recipes
	params := mux.Vars(r)

	db.DB.Find(&recipes, params["id"])

	if recipes.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("recipe not found"))
		return
	}

	db.DB.Unscoped().Delete(&recipes)
	w.WriteHeader(http.StatusNoContent)

}
