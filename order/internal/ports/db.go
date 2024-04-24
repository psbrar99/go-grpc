package ports

import "github.com/psbrar99/go-grpc/order/internal/application/core/domain"

type DbPort interface {
	GetOrder(id string) (domain.Order, error)
	SaveOrder(*domain.Order) error
}
