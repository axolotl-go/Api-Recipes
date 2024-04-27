package routes

import (
	"encoding/json"
	"net/http"

	"github.com/AxolotlJ-Dev/Api-Recipes/db"
	"github.com/AxolotlJ-Dev/Api-Recipes/models"
)

func GetUsersHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users"))
}

func GetUserHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}

func PostUserHandle(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}

func DeleteUserHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
