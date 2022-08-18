package routes

import (
	"golang/handlers"
	"golang/pkg/mysql"
	"golang/repositories"

	"github.com/gorilla/mux"
)

// Create UserRoutes function here ...
func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/carts", h.FindCarts).Methods("GET")
	r.HandleFunc("/cart/{id}", h.GetCart).Methods("GET")
	r.HandleFunc("/cart", h.CreateCart).Methods("POST")
	r.HandleFunc("/cart/{id}", h.UpdateCart).Methods("PATCH")
	r.HandleFunc("/cart/{id}", h.DeleteCart).Methods("DELETE")
}
