package domain

import "time"

type Orderitems struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int64   `json:"quantity"`
}

type Order struct {
	ID         int64        `json:"id"`
	CustomerID int64        `json:"customer_id"`
	Ordertime  int64        `json:"order_time"`
	OrderItems []Orderitems `json:"order_items"`
	Status     string       `json:"status"`
}

func NewOrder(o []Orderitems, customerid int64) Order {

	return Order{
		Ordertime:  time.Now().Unix(),
		CustomerID: customerid,
		OrderItems: o,
		Status:     "Pending",
	}

}
