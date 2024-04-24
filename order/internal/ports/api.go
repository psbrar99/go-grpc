package ports

import "github.com/psbrar99/go-grpc/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
