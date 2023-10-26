package interfaces

import "accountTransfer/domain"

type AccountRepo interface {
	Init()
	GetAll(page int, pageSize int) ([]domain.Account, error)
	UpdateOne(id string, account domain.Account) (domain.Account, error)
	FindById(id string) (domain.Account, error)
}
