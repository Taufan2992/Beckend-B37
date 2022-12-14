package routes

import (
	"golang/handlers"
	"golang/pkg/mysql"
	"golang/repositories"

	"github.com/gorilla/mux"
)

// Create UserRoutes function here ...
func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transaction", h.CreateTransaction).Methods("POST")
	r.HandleFunc("/transaction/{id}", h.UpdateTransaction).Methods("PATCH")
	r.HandleFunc("/transaction/{id}", h.DeleteTransaction).Methods("DELETE")
}
