package usecases

import (
	"accountTransfer/domain"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"sort"
)

type AccountRepoHandler struct {
	data map[string]domain.Account
}

func (u *AccountRepoHandler) Init() {

	file, err := os.Open("accounts-mock.json")

	if err != nil {
		log.Fatalln("Error opening the file:", err)
		return
	}

	data, err := io.ReadAll(file)

	if err != nil {
		log.Fatalln("Error reading the file:", err)
		return
	}

	var accounts []domain.Account

	if err := json.Unmarshal(data, &accounts); err != nil {
		log.Fatalln("Error unmarshaling JSON:", err)
		return
	}

	dataMap := make(map[string]domain.Account)

	for _, a := range accounts {

		dataMap[a.Id] = a
	}

	u.data = dataMap

	log.Print("Data file reading completed successfully")
}

func (u *AccountRepoHandler) GetAll(page int, pageSize int) ([]domain.Account, error) {

	start := (page - 1) * pageSize
	end := (page-1)*pageSize + pageSize

	var accounts []domain.Account

	for _, a := range u.data {

		accounts = append(accounts, a)
	}

	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i].Id < accounts[j].Id
	})

	if start >= len(accounts) || end >= len(accounts) {

		return nil, errors.New("index out of bounds")
	}

	return accounts[start:end], nil
}

func (u *AccountRepoHandler) FindById(id string) (domain.Account, error) {

	if _, ok := u.data[id]; !ok {
		return domain.Account{}, errors.New("ID not found")
	}

	return u.data[id], nil
}

func (u *AccountRepoHandler) UpdateOne(id string, account domain.Account) (domain.Account, error) {

	if _, ok := u.data[id]; !ok {
		return domain.Account{}, errors.New("ID not found")
	}

	u.data[id] = account

	return u.data[id], nil
}
