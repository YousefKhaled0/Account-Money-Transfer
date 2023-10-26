package main

import (
	"accountTransfer/controller"
	"accountTransfer/usecases"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	db := &usecases.AccountRepoHandler{}
	db.Init()

	interactor := &usecases.AccountInteractorHandler{
		DB: db,
	}

	accountController := &controller.AccountController{
		Interactor: interactor,
	}

	router := mux.NewRouter()
	router.HandleFunc("/transfer", accountController.TransferHandler).Methods("POST")
	router.HandleFunc("/accounts", accountController.GetAllAccountsHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
