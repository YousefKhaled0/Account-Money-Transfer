package controller_test

import (
	"accountTransfer/controller"
	"accountTransfer/usecases"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestTransferHandler(t *testing.T) {

	db := &usecases.AccountRepoHandler{}
	db.Init()

	accountController := &controller.AccountController{
		Interactor: &usecases.AccountInteractorHandler{
			DB: db,
		},
	}

	requestJSON := `{"from": "3d253e29-8785-464f-8fa0-9e4b57699db9", 
	"to": "17f904c1-806f-4252-9103-74e7a5d3e340", "amount": 10}`

	req, err := http.NewRequest("POST", "/transfer", strings.NewReader(requestJSON))

	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	accountController.TransferHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var response controller.TransferResponse

	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Errorf("Error decoding response JSON: %v", err)
	}

	if response.Balance != "77.11" {

		t.Errorf("Balance after transfer is not valid")
	}
}

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
