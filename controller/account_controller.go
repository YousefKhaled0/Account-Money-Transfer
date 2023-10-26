package controller

import (
	"accountTransfer/interfaces"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
)

type AccountController struct {
	Interactor interfaces.AccountInteractor
}

type transferRequest struct {
	FromID string  `json:"from" validate:"required"`
	ToID   string  `json:"to" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

type TransferResponse struct {
	FromAccount    string  `json:"fromAccount"`
	ToAccount      string  `json:"toAccount"`
	TransferAmount float64 `json:"transferAmount"`
	Balance        string  `json:"balanceAfterTransfer"`
}

func (accountController *AccountController) TransferHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	transferRequest := &transferRequest{}

	if err := decoder.Decode(&transferRequest); err != nil {
		log.Printf("Error decoding JSON request: %v", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	validate := validator.New()

	if err := validate.Struct(transferRequest); err != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	err := accountController.Interactor.Transfer(transferRequest.FromID,
		transferRequest.ToID, transferRequest.Amount)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fromAccount, _ := accountController.Interactor.GetOne(transferRequest.FromID)
	toAccount, _ := accountController.Interactor.GetOne(transferRequest.ToID)

	response, err := json.Marshal(&TransferResponse{
		fromAccount.Id, toAccount.Id, transferRequest.Amount, fromAccount.Balance,
	})

	if err != nil {
		http.Error(w, "Error marshaling JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (accountController *AccountController) GetAllAccountsHandler(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSizeStr := r.URL.Query().Get("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	accounts, err := accountController.Interactor.FindAll(page, pageSize)

	if err != nil {
		log.Printf("Error decoding JSON request: %v", err)
		http.Error(w, "Error fetching accounts", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(accounts)

	if err != nil {
		http.Error(w, "Error marshaling JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
