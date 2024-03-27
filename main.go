package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Food        string `json:"food"`
	Description string `json:"description"`
}

var albums = []album{

	{ID: "1", Title: "Quesadilla doble", Food: "Quesadilla", Description: "Una tortilla con queso en medio"},
	{ID: "2", Title: "Tacos de pollo", Food: "Tacos", Description: "Tortillas de maíz con pollo, cebolla y cilantro"},
	{ID: "3", Title: "Ensalada César", Food: "Ensalada", Description: "Lechuga romana, crutones, queso parmesano y aderezo César"},
	{ID: "4", Title: "Sopa de tomate", Food: "Sopa", Description: "Tomate, cebolla, ajo, caldo de pollo y hierbas"},
	{ID: "5", Title: "Pasta Alfredo", Food: "Pasta", Description: "Pasta con salsa cremosa de queso parmesano"},
	{ID: "6", Title: "Hamburguesa clásica", Food: "Hamburguesa", Description: "Carne de res, pan, lechuga, tomate, cebolla y queso"},
	{ID: "7", Title: "Pizza margarita", Food: "Pizza", Description: "Masa de pizza, salsa de tomate, mozzarella y albahaca"},
	{ID: "8", Title: "Pollo a la parrilla", Food: "Pollo", Description: "Pechuga de pollo marinada y cocida a la parrilla"},
	{ID: "9", Title: "Arroz frito", Food: "Arroz", Description: "Arroz cocido, huevo, cebolla, guisantes y salsa de soja"},
	{ID: "10", Title: "Sushi de salmón", Food: "Sushi", Description: "Arroz, salmón fresco, alga nori y salsa de soja"},
	{ID: "11", Title: "Ceviche de pescado", Food: "Ceviche", Description: "Pescado fresco marinado en limón con cebolla y cilantro"},
	{ID: "12", Title: "Tortilla española", Food: "Tortilla", Description: "Huevos, patatas, cebolla y aceite de oliva"},
	{ID: "13", Title: "Lasagna de carne", Food: "Lasagna", Description: "Capas de pasta, carne molida, queso y salsa de tomate"},
	{ID: "14", Title: "Empanadas argentinas", Food: "Empanadas", Description: "Masa rellena de carne, cebolla, huevo y especias"},
	{ID: "15", Title: "Sopa de fideos", Food: "Sopa", Description: "Caldo de pollo con fideos y verduras"},
	{ID: "16", Title: "Pastel de zanahoria", Food: "Pastel", Description: "Bizcocho de zanahoria con crema de queso"},
	{ID: "17", Title: "Costillas BBQ", Food: "Costillas", Description: "Costillas de cerdo cocidas lentamente en salsa BBQ"},
	{ID: "18", Title: "Gazpacho", Food: "Gazpacho", Description: "Sopa fría de tomate, pepino, pimiento, cebolla y ajo"},
	{ID: "19", Title: "Tiramisú", Food: "Postre", Description: "Capas de bizcocho, café, mascarpone y cacao en polvo"},
	{ID: "20", Title: "Camarones al ajillo", Food: "Camarones", Description: "Camarones cocidos en aceite de oliva con ajo y perejil"},
	{ID: "21", Title: "Ratatouille", Food: "Ratatouille", Description: "Guiso de verduras como berenjena, calabacín, pimiento y tomate"},
	{ID: "22", Title: "Guiso de lentejas", Food: "Guiso", Description: "Lentejas cocidas con cebolla, zanahoria, pimiento y chorizo"},
	{ID: "25", Title: "Crepas con Nutella", Food: "Crepas", Description: "Crepas rellenas de crema de avellanas Nutella"},
	{ID: "23", Title: "Filete de salmón a la parrilla", Food: "Salmón", Description: "Filete de salmón sazonado y cocido a la parrilla"},
	{ID: "24", Title: "Puré de patatas", Food: "Puré", Description: "Patatas cocidas y aplastadas con mantequilla y leche"},
	{ID: "25", Title: "Crepas con Nutella", Food: "Crepas", Description: "Crepas rellenas de crema de avellanas Nutella"},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	pegelagarto := "hola"
	fmt.Println(pegelagarto)
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:3000")
}
