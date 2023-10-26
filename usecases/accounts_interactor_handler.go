package usecases

import (
	"accountTransfer/domain"
	"accountTransfer/interfaces"
	"fmt"
	"strconv"
)

type AccountInteractorHandler struct {
	DB interfaces.AccountRepo
}

func (u *AccountInteractorHandler) FindAll(page int, pageSize int) ([]domain.Account, error) {

	if page < 1 || pageSize < 1 {
		return nil, fmt.Errorf("Invalid page or pageSize. Both must be greater than or equal to 1.")
	}

	return u.DB.GetAll(page, pageSize)
}

func (u *AccountInteractorHandler) GetOne(id string) (domain.Account, error) {

	account, err := u.DB.FindById(id)

	if err != nil {
		return domain.Account{}, err
	}

	return account, nil
}

func (u *AccountInteractorHandler) Transfer(fromId string, toId string, amount float64) error {

	fromAccount, err := u.DB.FindById(fromId)
	if err != nil {
		return err
	}

	toAccount, err := u.DB.FindById(toId)
	if err != nil {
		return err
	}

	fromBalance, err := strconv.ParseFloat(fromAccount.Balance, 64)
	if err != nil {
		return err
	}

	toBalance, err := strconv.ParseFloat(toAccount.Balance, 64)
	if err != nil {
		return err
	}

	if fromBalance < amount {
		return fmt.Errorf("Insufficient balance for transfer")
	}

	fromBalance -= amount
	toBalance += amount

	fromAccount.Balance = fmt.Sprintf("%.2f", fromBalance)
	toAccount.Balance = fmt.Sprintf("%.2f", toBalance)

	if _, err := u.DB.UpdateOne(fromId, fromAccount); err != nil {
		return err
	}

	if _, err := u.DB.UpdateOne(toId, toAccount); err != nil {
		return err
	}

	return nil
}
