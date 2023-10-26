package interfaces

import "accountTransfer/domain"

type AccountInteractor interface {
	FindAll(page int, pageSize int) ([]domain.Account, error)
	Transfer(fromId string, toId string, amount float64) error
	GetOne(id string) (domain.Account, error)
}
