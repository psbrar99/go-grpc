package grpc

import (
	"context"

	"github.com/psbrar99/go-grpc/order/internal/application/core/domain"
)

func (a Adapter) Create(ctx context.Context, request apis.CreateOrderRequest) (apis.CreateOrderResponse, error) {
	var orderItems []domain.Orderitems

	for _, oi := range request.OrderItems {
		orderItems = append(orderItems, domain.Orderitems{
			ProductCode: oi.ProductCode,
			UnitPrice:   oi.UnitPrice,
			Quantity:    oi.Quantity,
		})
	}
	newOrder := domain.NewOrder(orderItems, request.UserId)
	result, err := a.api.PlaceOrder(newOrder)
	if err != nil {
		return nil, err
	}

	return &apis.CreateOrderResponse{OrderID: result.ID}, nil

}
