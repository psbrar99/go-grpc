package db

import (
	"fmt"

	"github.com/psbrar99/go-grpc/order/internal/application/core/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID int64
	OrderItems []Orderitems
	Status     string
}

type Orderitems struct {
	gorm.Model
	ProductCode string
	UnitPrice   float64
	Quantity    int64
	OrderID     int64
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(url string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %v", err)
	}
	err = db.AutoMigrate(&Order{}, Orderitems{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil

}

func (a Adapter) GetOrder(id string) (domain.Order, error) {

}
