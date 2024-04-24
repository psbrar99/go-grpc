package api

import (
	"github.com/psbrar99/go-grpc/order/internal/application/core/domain"
	"github.com/psbrar99/go-grpc/order/internal/application/ports"
)

type Application struct {
	db ports.DBport
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {

	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil
}
